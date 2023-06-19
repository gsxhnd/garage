package crawl

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func (cc *crawlClient) getJavMovieInfoByJavbus(e *colly.HTMLElement) {
	var info = &JavMovie{}
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
}

func (cc *crawlClient) getJavMovieMagnetByJavbus(e *colly.HTMLElement) {
	fmt.Println(e)
}

func (cc *crawlClient) StartCrawlJavbusMovie(code string) error {
	cc.logger.Info("Download info: " + code)
	cc.collector.OnHTML("body script", cc.getJavMovieMagnetByJavbus)
	cc.collector.OnHTML(".container", cc.getJavMovieInfoByJavbus)
	cc.collector.OnRequest(func(r *colly.Request) {
		cc.logger.Info("Visiting" + r.URL.String())
	})
	if err := cc.collector.Visit(cc.javbusUrl + code); err != nil {
		return err
	}

	cc.collector.Wait()

	if len(cc.javInfos) == 0 {
		return nil
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

func (cc *crawlClient) StartCrawlJavbusMovieByPrefix() error {
	q, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})
	for i := cc.prefixMinNo; i <= cc.prefixMaxNo; i++ {
		code := fmt.Sprintf("%s-%03d", cc.prefixCode, i)
		q.AddURL(cc.javbusUrl + code)
	}
	cc.collector.OnRequest(func(r *colly.Request) {
		cc.logger.Info("Visiting" + r.URL.String())
	})
	cc.collector.OnHTML(".container", cc.getJavMovieInfoByJavbus)
	q.Run(cc.collector)
	cc.collector.Wait()

	cc.saveJavInfos()
	return nil
}

func (cc *crawlClient) StartCrawlJavbusMovieByStar(starCode string) error {
	return nil
}

func (cc *crawlClient) StartCrawlJavbusMovieByFilepath(inputPath string) error {
	q, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})
	var videoExt = []string{".avi", ".mp4", ".mkv"}
	if err := filepath.Walk(inputPath, func(path string, fi os.FileInfo, err error) error {
		if fi == nil {
			if err != nil {
				return err
			}
			return nil
		}
		if fi.IsDir() {
			return nil
		}
		filename := fi.Name()
		fileExt := filepath.Ext(filename)
		for _, b := range videoExt {
			if fileExt == b {
				q.AddURL(cc.javbusUrl + filename)
			}
		}
		return nil
	}); err != nil {
		return err
	}

	cc.collector.OnRequest(func(r *colly.Request) {
		cc.logger.Info("Visiting" + r.URL.String())
	})
	cc.collector.OnHTML(".container", cc.getJavMovieInfoByJavbus)
	q.Run(cc.collector)
	cc.collector.Wait()
	return nil
}
