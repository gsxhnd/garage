package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

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
