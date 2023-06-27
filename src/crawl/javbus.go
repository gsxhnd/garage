package crawl

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/inhies/go-bytesize"
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
	e.ForEach("script", func(i int, element *colly.HTMLElement) {
		if i != 2 {
			return
		}
		param := element.Text
		param = strings.Replace(param, " ", "", -1)
		param = strings.Replace(param, "\tvar", "", -1)
		param = strings.Replace(param, "\r", "", -1)
		param = strings.Replace(param, "\r\n", "", -1)
		param = strings.Replace(param, "\n", "", -1)
		param = strings.Replace(param, ";", "&", -1)
		param = strings.Replace(param, "'", "", -1)
		urlS := "https://www.javbus.com/ajax/uncledatoolsbyajax.php?" + param + "lang=zh&floor=442"
		cc.logger.Info("Get magnet url: " + urlS)

		r, _ := http.NewRequest("GET", urlS, nil)
		r.Header.Add("Referer", e.Request.URL.Scheme+"://"+e.Request.URL.Host+e.Request.URL.Path)
		res, err := cc.httpClient.Do(r)
		if err != nil {
			cc.logger.Error("http client error: " + err.Error())
			return
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			cc.logger.Error("http read response error: " + err.Error())
			return
		}
		defer res.Body.Close()

		doc, err := htmlquery.Parse(strings.NewReader("<table><tbody>" + string(body) + "</tbody></table>"))
		if err != nil {
			cc.logger.Error("html query error: " + err.Error())
		}
		list, err := htmlquery.QueryAll(doc, "//tr")
		if err != nil {
			cc.logger.Error("html query tr error: " + err.Error())
		}
		if len(list) == 0 {
			cc.logger.Info("当前无磁力连接")
			return
		}

		var mList = make([]JavMovieMagnet, 0)
		for _, n := range list {
			tdList, _ := htmlquery.QueryAll(n, "//td/a")
			var m = JavMovieMagnet{
				HD:       false,
				Subtitle: false,
			}
			for tdIndex, tdValue := range tdList {
				switch tdIndex {
				case 0:
					m.Link = htmlquery.SelectAttr(tdValue, "href")
					m.Name = htmlquery.InnerText(tdValue)
				default:
					if htmlquery.InnerText(tdValue) == "高清" {
						m.HD = true
					} else if htmlquery.InnerText(tdValue) == "字幕" {
						m.Subtitle = true
					} else {
						var sizeStr string = htmlquery.InnerText(tdValue)
						sizeStr = strings.Replace(sizeStr, " ", "", -1)
						sizeStr = strings.Replace(sizeStr, "\n", "", -1)
						sizeStr = strings.Replace(sizeStr, "\x09", "", -1)
						_, err := time.Parse("2006-01-02", sizeStr)
						if err != nil {
							b, err := bytesize.Parse(sizeStr)
							if err != nil {
								return
							}
							sizeStr = strings.Replace(b.Format("%.2f", "MB", false), "MB", "", -1)
							size, _ := strconv.ParseFloat(sizeStr, 64)
							m.Size = size
						}
					}
				}
			}
			mList = append(mList, m)
		}

		var maxSize float64 = 0
		var bestMagnet string = ""
		for _, m := range mList {
			if m.Size > maxSize {
				maxSize = m.Size
				bestMagnet = m.Link
			}
		}
		cc.javMagnets = append(cc.javMagnets, bestMagnet)
	})
}

func (cc *crawlClient) StartCrawlJavbusMovie(code string) error {
	cc.logger.Info("Download info: " + code)
	cc.collector.OnRequest(func(r *colly.Request) {
		cc.logger.Info("Visiting: " + r.URL.String())
	})

	if cc.downloadMagent {
		cc.collector.OnHTML("body", cc.getJavMovieMagnetByJavbus)
	}
	cc.collector.OnHTML(".container", cc.getJavMovieInfoByJavbus)

	if err := cc.collector.Visit(cc.javbusUrl + "/" + code); err != nil {
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
	cc.saveMagents()
	return nil
}

func (cc *crawlClient) StartCrawlJavbusMovieByPrefix() error {
	q, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})

	for i := cc.prefixMinNo; i <= cc.prefixMaxNo; i++ {
		code := fmt.Sprintf("%s-%03d", cc.prefixCode, i)
		q.AddURL(cc.javbusUrl + "/" + code)
	}

	cc.collector.OnRequest(func(r *colly.Request) {
		cc.logger.Info("Visiting " + r.URL.String())
	})
	if cc.downloadMagent {
		cc.collector.OnHTML("body", cc.getJavMovieMagnetByJavbus)
	}

	cc.collector.OnHTML(".container", cc.getJavMovieInfoByJavbus)

	q.Run(cc.collector)
	cc.collector.Wait()

	cc.saveJavInfos()
	for _, v := range cc.javInfos {
		err := cc.saveCovers(v.Cover, v.Code)
		if err != nil {
			return err
		}
	}
	cc.saveMagents()
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
				filePrefix := strings.Replace(filename, b, "", -1)
				q.AddURL(cc.javbusUrl + "/" + filePrefix)
			}
		}
		return nil
	}); err != nil {
		return err
	}

	cc.collector.OnRequest(func(r *colly.Request) {
		cc.logger.Info("Visiting: " + r.URL.String())
	})
	cc.collector.OnHTML(".container", cc.getJavMovieInfoByJavbus)

	q.Run(cc.collector)
	cc.collector.Wait()

	cc.saveJavInfos()
	for _, v := range cc.javInfos {
		err := cc.saveCovers(v.Cover, v.Code)
		if err != nil {
			return err
		}
	}
	cc.saveMagents()
	return nil
}
