package garage

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type Client struct {
	collector *colly.Collector
	proxy     string
	maxDepth  int
	javbusUrl string
	javlibUrl string
}

type JavMovie struct {
	Code           string `json:"code"`
	Title          string `json:"title"`
	ProduceCompany string `json:"produce_company"`
	PublishCompany string `json:"publish_company"`
}

func NewCrawlClient() *Client {
	return &Client{
		collector: colly.NewCollector(),
		proxy:     "",
		maxDepth:  100,
		javbusUrl: "https://www.javbus.com/",
		javlibUrl: "https://www.javbus.com/",
	}
}

func (c2 *Client) SetProxy(proxy string) (err error) {
	err = c2.collector.SetProxy(proxy)
	return
}

func (c2 *Client) StarCrawlJavbusMovie(code string) {
	c2.collector.OnHTML(".container", func(e *colly.HTMLElement) {
		var data JavMovie
		data.Title = e.ChildText("h3")
		cover := e.ChildAttr(".screencap img", "src")
		e.ForEach(".info p", func(i int, element *colly.HTMLElement) {
			switch i {
			case 0:
				data.Code = element.ChildTexts("span")[1]
			case 1:
				//jav.PublishDate = element.Text
			case 3:
				data.ProduceCompany = element.ChildText("a")
			case 4:
				data.PublishCompany = element.ChildText("a")
			}
		})
		fmt.Println(data)
		fmt.Println(cover)
	})
	_ = c2.collector.Visit(c2.javbusUrl + code)
}
