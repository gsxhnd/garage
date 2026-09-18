package jav

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Source identifies a crawl site.
type Source string

const (
	SourceJavbus Source = "javbus"
	SourceJavdb  Source = "javdb"

	defaultJavbusURL = "https://www.javbus.com"
	defaultJavdbURL  = "https://www.javdb.com"
	defaultTimeout   = 30 * time.Second
	maxListPages     = 200
)

var (
	errNotFound   = errors.New("jav: movie not found")
	errEmptyQuery = errors.New("jav: empty query")
)

// Crawler fetches movie metadata (and optionally magnets) from a site.
type Crawler interface {
	Source() Source
	Crawl(ctx context.Context, query Query) ([]Movie, error)
}

// Logger is the subset of structured logging used by crawlers and the store.
type Logger interface {
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
}

// Config controls HTTP, output, and optional downloads.
type Config struct {
	Proxy          string
	Cookie         string
	BaseURL        string
	OutDir         string
	DownloadMagnet bool
	DownloadCover  bool
	RandomDelay    time.Duration
	Parallelism    int
	Timeout        time.Duration
}

// Query describes what to crawl. Codes, prefix expansion, a videos directory,
// and actor/star IDs can be combined in one request.
type Query struct {
	Codes      []string
	Prefixes   []string
	PrefixMin  uint64
	PrefixMax  uint64
	PrefixZero uint64
	StarIDs    []string
	VideosDir  string
}

// NewCrawler returns a site-specific Crawler.
func NewCrawler(source Source, logger Logger, cfg Config) (Crawler, error) {
	if logger == nil {
		return nil, fmt.Errorf("jav: logger is nil")
	}
	cfg, err := cfg.withDefaults(source)
	if err != nil {
		return nil, err
	}
	httpClient, err := newHTTPClient(cfg)
	if err != nil {
		return nil, err
	}
	eng := &engine{logger: logger, cfg: cfg, http: httpClient}
	switch source {
	case SourceJavbus:
		return &javbusCrawler{engine: eng}, nil
	case SourceJavdb:
		return &javdbCrawler{engine: eng}, nil
	default:
		return nil, fmt.Errorf("jav: unsupported source %q", source)
	}
}

func (c Config) withDefaults(source Source) (Config, error) {
	c.applyCommonDefaults()
	if c.BaseURL == "" {
		switch source {
		case SourceJavbus:
			c.BaseURL = defaultJavbusURL
		case SourceJavdb:
			c.BaseURL = defaultJavdbURL
		default:
			return c, fmt.Errorf("jav: unsupported source %q", source)
		}
	}
	c.BaseURL = trimSlash(c.BaseURL)
	return c, nil
}

func (c *Config) applyCommonDefaults() {
	if c.Parallelism <= 0 {
		c.Parallelism = 1
	}
	if c.Timeout <= 0 {
		c.Timeout = defaultTimeout
	}
	if c.OutDir == "" {
		c.OutDir = "."
	}
}

func (q Query) resolvedCodes() ([]string, error) {
	var codes []string
	codes = append(codes, q.Codes...)
	if len(q.Prefixes) > 0 {
		if q.PrefixMax < q.PrefixMin {
			return nil, fmt.Errorf("jav: prefix_max < prefix_min")
		}
		for _, prefix := range q.Prefixes {
			codes = append(codes, ExpandPrefix(prefix, q.PrefixMin, q.PrefixMax, q.PrefixZero)...)
		}
	}
	if q.VideosDir != "" {
		fromDir, err := CodesFromVideosDir(q.VideosDir)
		if err != nil {
			return nil, err
		}
		codes = append(codes, fromDir...)
	}
	return uniquePreserve(codes), nil
}

type engine struct {
	logger Logger
	cfg    Config
	http   *httpClient
}

func (e *engine) crawlTargets(ctx context.Context, targets []string, fetch func(context.Context, string) (Movie, error)) ([]Movie, error) {
	if len(targets) == 0 {
		return nil, nil
	}
	if e.cfg.Parallelism <= 1 {
		return e.crawlSequential(ctx, targets, fetch)
	}
	return e.crawlParallel(ctx, targets, fetch)
}

func (e *engine) crawlSequential(ctx context.Context, targets []string, fetch func(context.Context, string) (Movie, error)) ([]Movie, error) {
	movies := make([]Movie, 0, len(targets))
	var nFail int
	for _, target := range targets {
		if err := ctx.Err(); err != nil {
			return movies, err
		}
		movie, err := fetch(ctx, target)
		if err != nil {
			if errors.Is(err, errNotFound) {
				e.logger.Warnw("movie not found", "target", target)
				continue
			}
			nFail++
			e.logger.Errorw("crawl movie failed", "target", target, "error", err)
			continue
		}
		movies = append(movies, movie)
	}
	if len(movies) == 0 && nFail > 0 {
		return nil, fmt.Errorf("jav: failed to crawl %d targets", nFail)
	}
	return movies, nil
}

func (e *engine) crawlParallel(ctx context.Context, targets []string, fetch func(context.Context, string) (Movie, error)) ([]Movie, error) {
	var (
		mu     sync.Mutex
		wg     sync.WaitGroup
		movies = make([]Movie, 0, len(targets))
		nFail  int
	)
	sem := make(chan struct{}, e.cfg.Parallelism)
	for _, target := range targets {
		target := target
		select {
		case <-ctx.Done():
			wg.Wait()
			return movies, ctx.Err()
		case sem <- struct{}{}:
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			movie, err := fetch(ctx, target)
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				if errors.Is(err, errNotFound) {
					e.logger.Warnw("movie not found", "target", target)
					return
				}
				nFail++
				e.logger.Errorw("crawl movie failed", "target", target, "error", err)
				return
			}
			movies = append(movies, movie)
		}()
	}
	wg.Wait()
	if len(movies) == 0 && nFail > 0 {
		return nil, fmt.Errorf("jav: failed to crawl %d targets", nFail)
	}
	return movies, nil
}

func uniquePreserve(in []string) []string {
	seen := make(map[string]struct{}, len(in))
	out := make([]string, 0, len(in))
	for _, s := range in {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	return out
}

func trimSlash(s string) string {
	for len(s) > 1 && s[len(s)-1] == '/' {
		s = s[:len(s)-1]
	}
	return s
}
