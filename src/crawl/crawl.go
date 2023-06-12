package crawl

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

type CrawlClient interface {
	StartCrawlJavbusMovie(code string) error
	StartCrawlJavbusMovieByPrefix() error
	StartCrawlJavbusMovieByStar(starCode string) error
	setProxy(proxy string) error
	mkAllDir() error
	// getJavMovieInfoByJavbus(code string) error
	getJavMovieInfoByJavbus(element *colly.HTMLElement)
	saveJavInfos() error
	saveCovers(coverPath, name string) error
	saveMagents() error
}
type crawlClient struct {
	logger      *zap.Logger
	collector   *colly.Collector
	httpClient  *http.Client
	maxDepth    int
	javbusUrl   string
	javlibUrl   string
	javInfos    []JavMovie
	destPath    string
	prefixCode  string
	prefixMinNo int
	prefixMaxNo int
}

type CrawlOptions struct {
	DestPath    string
	Proxy       string
	PrefixCode  string
	PrefixMinNo int
	PrefixMaxNo int
}

func NewCrawlClient(logger *zap.Logger, option CrawlOptions) (CrawlClient, error) {
	var client = &crawlClient{
		collector:   colly.NewCollector(),
		httpClient:  &http.Client{},
		maxDepth:    100,
		javbusUrl:   "https://www.javbus.com/",
		javlibUrl:   "https://www.javbus.com/",
		logger:      logger,
		javInfos:    make([]JavMovie, 0),
		destPath:    option.DestPath,
		prefixCode:  option.PrefixCode,
		prefixMinNo: option.PrefixMinNo,
		prefixMaxNo: option.PrefixMaxNo,
	}
	if err := client.setProxy(option.Proxy); err != nil {
		return nil, err
	}

	if err := client.mkAllDir(); err != nil {
		return nil, err
	} else {
		return client, nil
	}
}

func (cc *crawlClient) setProxy(proxy string) error {
	if err := cc.collector.SetProxy(proxy); err != nil {
		return err
	}
	uri, err := url.Parse(proxy)
	if err != nil {
		return err
	}
	cc.httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},
	}
	return nil
}

func (cc *crawlClient) mkAllDir() error {
	fullPath := filepath.Join(cc.destPath, "cover")
	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(fullPath, os.ModePerm)
			if err != nil {
				return err
			}
		} else if os.IsExist(err) {
			return nil
		} else {
			return err
		}
	}
	return nil
}
