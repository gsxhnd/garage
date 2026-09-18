package jav

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/gsxhnd/garage/utils"
)

// Store persists crawl results.
type Store interface {
	Save(ctx context.Context, movies []Movie) error
}

var _ Store = (*localStore)(nil)

type localStore struct {
	logger Logger
	cfg    Config
	http   *httpClient
}

// NewLocalStore writes CSV, magnet lists, and optional covers under cfg.OutDir.
func NewLocalStore(logger Logger, cfg Config) (Store, error) {
	if logger == nil {
		return nil, fmt.Errorf("jav: logger is nil")
	}
	cfg.applyCommonDefaults()
	httpClient, err := newHTTPClient(cfg)
	if err != nil {
		return nil, err
	}
	if err := utils.MakeDir(cfg.OutDir); err != nil {
		return nil, err
	}
	if cfg.DownloadCover {
		if err := utils.MakeDir(filepath.Join(cfg.OutDir, "cover")); err != nil {
			return nil, err
		}
	}
	return &localStore{logger: logger, cfg: cfg, http: httpClient}, nil
}

func (s *localStore) Save(ctx context.Context, movies []Movie) error {
	if len(movies) == 0 {
		s.logger.Warnw("no jav info to save")
		return nil
	}

	stamp := time.Now().Local().Format("2006-01-02-15-04-05")
	csvPath := filepath.Join(s.cfg.OutDir, stamp+"-jav_info.csv")
	if err := s.writeCSV(csvPath, movies); err != nil {
		return err
	}

	if hasMagnets(movies) {
		magnetPath := filepath.Join(s.cfg.OutDir, stamp+"-jav_magnet.txt")
		if err := s.writeMagnets(magnetPath, movies); err != nil {
			return err
		}
	}

	if s.cfg.DownloadCover {
		s.downloadCovers(ctx, movies)
	}
	return nil
}

func (s *localStore) writeCSV(path string, movies []Movie) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		s.logger.Errorw("save jav info file failed", "error", err)
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	header := []string{
		"code", "title", "cover", "publish_date", "length",
		"director", "produce_company", "publish_company", "series", "stars", "page_url",
	}
	if err := w.Write(header); err != nil {
		return err
	}
	for _, m := range movies {
		row := []string{
			m.Code, m.Title, m.Cover, m.PublishDate, m.Length,
			m.Director, m.ProduceCompany, m.PublishCompany, m.Series, m.Stars, m.PageURL,
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}

func (s *localStore) writeMagnets(path string, movies []Movie) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		s.logger.Errorw("save jav magnet file failed", "error", err)
		return err
	}
	defer f.Close()

	for _, m := range movies {
		for _, mag := range m.Magnets {
			if mag.Link == "" {
				continue
			}
			if _, err := fmt.Fprintf(f, "%s\t%s\n", m.Code, mag.Link); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *localStore) downloadCovers(ctx context.Context, movies []Movie) {
	for _, m := range movies {
		if err := ctx.Err(); err != nil {
			return
		}
		if m.Cover == "" {
			continue
		}
		u, err := url.Parse(m.Cover)
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
			s.logger.Warnw("skip invalid cover url", "code", m.Code, "cover", m.Cover)
			continue
		}
		body, err := s.http.Get(ctx, m.Cover, nil)
		if err != nil {
			s.logger.Errorw("download cover failed", "code", m.Code, "url", m.Cover, "error", err)
			continue
		}
		ext := path.Ext(u.Path)
		if ext == "" {
			ext = ".jpg"
		}
		out := filepath.Join(s.cfg.OutDir, "cover", safeFileName(m.Code)+ext)
		if err := os.WriteFile(out, body, 0o644); err != nil {
			s.logger.Errorw("write cover failed", "path", out, "error", err)
		}
	}
}

func hasMagnets(movies []Movie) bool {
	for _, m := range movies {
		for _, mag := range m.Magnets {
			if mag.Link != "" {
				return true
			}
		}
	}
	return false
}

func safeFileName(code string) string {
	code = strings.TrimSpace(code)
	code = strings.ReplaceAll(code, "/", "-")
	code = strings.ReplaceAll(code, "\\", "-")
	if code == "" {
		return "unknown"
	}
	return code
}
