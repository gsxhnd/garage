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
	// 通过番号爬取对应的电影信息
	StartCrawlJavbusMovie(code string) error
	// 通过番号前缀爬取对应的电影信息
	StartCrawlJavbusMovieByPrefix() error
	// 通过演员ID爬取对应的电影信息
	StartCrawlJavbusMovieByStar(starCode string) error
	// 访问文件夹下的视频列表爬取电影信息
	StartCrawlJavbusMovieByFilepath(inputPath string) error
	// 设置代理
	setProxy(proxy string) error
	// 创建输出目录
	mkAllDir() error
	getJavMovieInfoByJavbus(element *colly.HTMLElement)
	getJavMovieMagnetByJavbus(e *colly.HTMLElement)
	// 保存CSV格式的电影信息
	saveJavInfos() error
	// 下载电影封面
	saveCovers(coverPath, name string) error
	// 下载磁力列表
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
