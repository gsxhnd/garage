package crawl

import (
	"net/http"
	"net/url"

	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

type CrawlClient interface {
	SetProxy(proxy string) (err error)
	StartCrawlJavbusMovie(code string) error
	StartCrawlJavbusMovieByPrefix(prefixCode string) error
	StartCrawlJavbusMovieByStar(starCode string) error
	saveJavInfos() error
	saveCovers() error
	saveMagents() error
}
type crawlClient struct {
	logger     *zap.Logger
	collector  *colly.Collector
	httpClient *http.Client
	maxDepth   int
	javbusUrl  string
	javlibUrl  string
	javInfos   []JavMovie
}

func NewCrawlClient(logger *zap.Logger) CrawlClient {
	return &crawlClient{
		collector:  colly.NewCollector(),
		httpClient: &http.Client{},
		maxDepth:   100,
		javbusUrl:  "https://www.javbus.com/",
		javlibUrl:  "https://www.javbus.com/",
		logger:     logger,
		javInfos:   []JavMovie{},
	}
}

func (cc *crawlClient) SetProxy(proxy string) (err error) {
	uri, _ := url.Parse(proxy)
	err = cc.collector.SetProxy(proxy)
	cc.httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},
	}
	return
}
