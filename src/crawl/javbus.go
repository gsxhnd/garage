package crawl

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

func (cc *crawlClient) getJavMovieInfoByJavbus(code string) error {
	var info = &JavMovie{}
	cc.collector.OnHTML(".container", func(e *colly.HTMLElement) {
		info.Title = e.ChildText("h3")
		info.Cover = e.ChildAttr(".screencap img", "src")
		e.ForEach(".info p", func(i int, element *colly.HTMLElement) {
			key := element.ChildText("span")
			switch i {
			case 0:
				info.Code = element.ChildTexts("span")[1]
			}
			switch key {
			case "發行日期:":
				pd := element.Text
				info.PublishDate = strings.Split(pd, " ")[1]
			case "長度:":
				pd := element.Text
				p := strings.Split(pd, " ")[1]
				info.Length = strings.Split(p, "分鐘")[0]
			case "導演:":
				info.Director = element.ChildText("a")
			case "製作商:":
				info.ProduceCompany = element.ChildText("a")
			case "發行商:":
				info.PublishCompany = element.ChildText("a")
			case "系列:":
				info.Series = element.ChildText("a")
			}
		})
		e.ForEach("ul li .star-name a", func(i int, element *colly.HTMLElement) {
			star := element.Attr("title")
			info.Stars += star + ";"
		})
		cc.javInfos = append(cc.javInfos, *info)
	})
	err := cc.collector.Visit(cc.javbusUrl + code)
	return err
}

func (cc *crawlClient) StartCrawlJavbusMovie(code string) error {
	cc.logger.Info("Download info: " + code)
	if err := cc.getJavMovieInfoByJavbus(code); err != nil {
		return err
	}
	if err := cc.saveJavInfos(); err != nil {
		return err
	}
	for _, v := range cc.javInfos {
		err := cc.saveCovers(v.Cover, v.Code)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cc *crawlClient) StartCrawlJavbusMovieByPrefix(prefixCode string) error {
	return nil
}

func (cc *crawlClient) StartCrawlJavbusMovieByStar(starCode string) error {
	return nil
}
