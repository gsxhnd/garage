package garage_jav

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

type JavCrawl struct {
	javInfos []JavMovie
}

func NewJavCrawl() *JavCrawl {
	return &JavCrawl{
		javInfos: make([]JavMovie, 0),
	}
}

func (jc *JavCrawl) GetJavMovieInfowByJavbus(e *colly.HTMLElement) {
	var Infow = &JavMovie{}
	Infow.Title = e.ChildText("h3")
	Infow.Cover = e.ChildAttr(".screencap img", "src")
	e.ForEach(".Infow p", func(i int, element *colly.HTMLElement) {
		key := element.ChildText("span")
		switch i {
		case 0:
			Infow.Code = element.ChildTexts("span")[1]
		}
		switch key {
		case "發行日期:":
			pd := element.Text
			Infow.PublishDate = strings.Split(pd, " ")[1]
		case "長度:":
			pd := element.Text
			p := strings.Split(pd, " ")[1]
			Infow.Length = strings.Split(p, "分鐘")[0]
		case "導演:":
			Infow.Director = element.ChildText("a")
		case "製作商:":
			Infow.ProduceCompany = element.ChildText("a")
		case "發行商:":
			Infow.PublishCompany = element.ChildText("a")
		case "系列:":
			Infow.Series = element.ChildText("a")
		}
	})
	e.ForEach("ul li .star-name a", func(i int, element *colly.HTMLElement) {
		star := element.Attr("title")
		Infow.Stars += star + ";"
	})
	jc.javInfos = append(jc.javInfos, *Infow)
}


func (jc *JavCrawl) GetJavStarMovieByJavbus(e *colly.HTMLElement) {
	// e.ForEach("#waterfall > div", func(i int, element *colly.HTMLElement) {
	// 	cc.collectorQueue.AddURL(element.ChildAttr("a", "href"))
	// })
	// e.ForEach("div.text-center.hidden-xs > ul", func(i int, element *colly.HTMLElement) {
	// 	page := element.ChildAttr("a#next", "href")
	// 	if page != "" {
	// 		cc.pageCollector.Visit(cc.javbusUrl + element.ChildAttr("a#next", "href"))
	// 	}
	// })
}