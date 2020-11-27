package core

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func GetSubject() {
	var c = colly.NewCollector()
	c.OnHTML("html", func(e *colly.HTMLElement) {

		// Print link
		fmt.Println("on html:")
		fmt.Println(e.Attr("meta"))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("res:", string(response.Body))
	})

	_ = c.Visit("https://www.baidu.com")
}
