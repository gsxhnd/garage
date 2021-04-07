package crawl

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

var (
	c *colly.Collector
)

func init() {
	c = colly.NewCollector()
	c.MaxDepth = 100
}

func SetProxy(proxy string) {
	_ = c.SetProxy(proxy)
}

func StartCrawl(url string) {
	c.OnHTML(".container-fluid", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})
	_ = c.Visit(url)
}
