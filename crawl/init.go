package crawl

import (
	"fmt"
	"garage/model"
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
	c.OnHTML(".container", func(e *colly.HTMLElement) {
		var jav = new(model.JavMovie)

		jav.Title = e.ChildText("h3")
		cover := e.ChildAttr(".screencap img", "src")
		e.ForEach(".info p", func(i int, element *colly.HTMLElement) {
			//fmt.Println(i)
			//fmt.Println(element)
			switch i {
			case 0:
				jav.Code = element.ChildTexts("span")[1]
			case 1:
				//jav.PublishDate = element.Text
			case 3:
				jav.ProduceCompany = element.ChildText("a")
			case 4:
				jav.PublishCompany = element.ChildText("a")
			}
		})
		fmt.Println(jav)
		fmt.Println(cover)
	})
	_ = c.Visit(url)
}
