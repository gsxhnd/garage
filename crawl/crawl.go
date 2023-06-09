package crawl

import (
	"net/http"
	"net/url"

	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

type CrawlClient struct {
	collector  *colly.Collector
	httpClient *http.Client
	maxDepth   int
	javbusUrl  string
	javlibUrl  string
	logger     *zap.Logger
}

type JavMovie struct {
	Code           string   `json:"code"`
	Title          string   `json:"title"`
	Cover          string   `json:"cover"`
	PublishDate    string   `json:"publish_date"`
	Length         string   `json:"length"`
	Director       string   `json:"director"`
	ProduceCompany string   `json:"produce_company"`
	PublishCompany string   `json:"publish_company"`
	Series         string   `json:"series"`
	Stars          []string `json:"stars"`
}

type JavMovieMagnet struct {
	Link string `json:"link"`
	Size string `json:"size"`
}

func NewCrawlClient(logger *zap.Logger) *CrawlClient {
	return &CrawlClient{
		collector:  colly.NewCollector(),
		httpClient: &http.Client{},
		maxDepth:   100,
		javbusUrl:  "https://www.javbus.com/",
		javlibUrl:  "https://www.javbus.com/",
		logger:     logger,
	}
}

func (cc *CrawlClient) SetProxy(proxy string) (err error) {
	uri, _ := url.Parse(proxy)
	err = cc.collector.SetProxy(proxy)
	cc.httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},
	}
	return
}
