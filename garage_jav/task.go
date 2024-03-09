package garage_jav

import (
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/inhies/go-bytesize"
	"github.com/reactivex/rxgo/v2"
)

type task struct {
	id        string
	name      string
	data      []JavMovie
	collector *colly.Collector
	http      *resty.Client
	ob        rxgo.Observable
	javMovies []JavMovie
	javbusUrl string
}

func NewJavbusTask(collector *colly.Collector, http *resty.Client) *task {
	return &task{
		id:        "",
		name:      "",
		data:      make([]JavMovie, 0),
		collector: collector,
		http:      http,
		ob:        rxgo.Empty(),
		javbusUrl: "https://www.javbus.com",
	}
}

func (t *task) getTaskInfo() {}

func (t *task) getJavbusMoviesInfo(codes []string) {
	queue, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})
	for _, code := range codes {
		queue.AddURL(t.javbusUrl + "/" + code)
	}
	t.collector.OnHTML(".container", t.javbusMovieInfoCrawl)
	queue.Run(t.collector)
}

func (t *task) getJavbusMovieByStar() {
	collector := t.collector.Clone()
	pageCollecctor := t.collector.Clone()
	queue, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})

	pageCollecctor.OnHTML("body", func(h *colly.HTMLElement) {
		h.ForEach("#waterfall > div", func(i int, element *colly.HTMLElement) {
			queue.AddURL(element.ChildAttr("a", "href"))
		})
		h.ForEach("div.text-center.hidden-xs > ul", func(i int, element *colly.HTMLElement) {
			page := element.ChildAttr("a#next", "href")
			if page != "" {
				pageCollecctor.Visit(t.javbusUrl + element.ChildAttr("a#next", "href"))
			}
		})
	})

	collector.OnHTML(".container", t.javbusMovieInfoCrawl)

	queue.Run(collector)
}

func (t *task) getJavMovieMagnetByJavbus(e *colly.HTMLElement) {
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
		// t.logger.Infow("Get magnet url: " + urlS)

		res, err := t.http.R().SetHeaders(map[string]string{
			"Referer": e.Request.URL.Scheme + "://" + e.Request.URL.Host + e.Request.URL.Path,
		}).Get(urlS)
		if err != nil {
			// cc.logger.Errorw("http client error: " + err.Error())
			return
		}

		body, err := io.ReadAll(res.RawResponse.Body)
		if err != nil {
			// cc.logger.Errorw("http read response error: " + err.Error())
			return
		}
		// defer res.Body.Close()

		doc, err := htmlquery.Parse(strings.NewReader("<table><tbody>" + string(body) + "</tbody></table>"))
		if err != nil {
			// cc.logger.Errorw("html query error: " + err.Error())
		}
		list, err := htmlquery.QueryAll(doc, "//tr")
		if err != nil {
			// cc.logger.Errorw("html query tr error: " + err.Error())
		}
		if len(list) == 0 {
			// cc.logger.Infow("当前无磁力连接")
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
		// var bestMagnet string = ""
		for _, m := range mList {
			if m.Size > maxSize {
				maxSize = m.Size
				// bestMagnet = m.Link
			}
		}

		// TODO: jav code need
		t.javMovies[0].Magnets = append(t.javMovies[0].Magnets, mList...)
	})
}

func (t *task) javbusMovieInfoCrawl(e *colly.HTMLElement) {
	var Info = &JavMovie{}
	Info.Title = e.ChildText("h3")
	Info.Cover = e.ChildAttr(".screencap img", "src")
	e.ForEach(".Info p", func(i int, element *colly.HTMLElement) {
		key := element.ChildText("span")
		switch i {
		case 0:
			Info.Code = element.ChildTexts("span")[1]
		}
		switch key {
		case "發行日期:":
			pd := element.Text
			Info.PublishDate = strings.Split(pd, " ")[1]
		case "長度:":
			pd := element.Text
			p := strings.Split(pd, " ")[1]
			Info.Length = strings.Split(p, "分鐘")[0]
		case "導演:":
			Info.Director = element.ChildText("a")
		case "製作商:":
			Info.ProduceCompany = element.ChildText("a")
		case "發行商:":
			Info.PublishCompany = element.ChildText("a")
		case "系列:":
			Info.Series = element.ChildText("a")
		}
	})

	e.ForEach("ul li .star-name a", func(i int, element *colly.HTMLElement) {
		star := element.Attr("title")
		Info.Stars += star + ";"
	})

	t.javMovies = append(t.javMovies, *Info)
}
