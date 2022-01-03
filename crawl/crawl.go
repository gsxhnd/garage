package crawl

import (
	"net/http"
	"net/url"

	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

type Client struct {
	collector  *colly.Collector
	httpClient *http.Client
	proxy      string
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

func NewCrawlClient(logger *zap.Logger) *Client {
	return &Client{
		collector:  colly.NewCollector(),
		httpClient: &http.Client{},
		proxy:      "",
		maxDepth:   100,
		javbusUrl:  "https://www.javbus.com/",
		javlibUrl:  "https://www.javbus.com/",
		logger:     logger,
	}
}

func (c2 *Client) SetProxy(proxy string) (err error) {
	uri, _ := url.Parse(proxy)
	err = c2.collector.SetProxy(proxy)
	c2.httpClient = &http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(uri),
		},
	}
	return
}
