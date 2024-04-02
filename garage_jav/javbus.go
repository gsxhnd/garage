package garage_jav

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-resty/resty/v2"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/gsxhnd/garage/utils"
	"github.com/inhies/go-bytesize"
)

const JAVBUS_URL = "https://www.javbus.com"

type JavbusCrawl interface {
	GetJavbusMovieByHomePage() ([]JavMovie, error) // 通过首页爬取对应的电影信息
	GetJavbusMovie() ([]JavMovie, error)           // 通过番号爬取对应的电影信息
	GetJavbusMovieByPrefix() ([]JavMovie, error)   // 通过番号前缀爬取对应的电影信息
	GetJavbusMovieByStar() ([]JavMovie, error)     // 通过演员ID爬取对应的电影信息
	GetJavbusMovieByFilepath() ([]JavMovie, error) // 访问文件夹下的视频列表爬取电影信息
	SaveLocal(infos []JavMovie) error
}

type javbusCrawl struct {
	logger     utils.Logger
	opt        *JavbusOption
	collector  *colly.Collector
	httpClient *resty.Client
}

func NewJavbusCrawl(logger utils.Logger, opt *JavbusOption, config *CrawlConfig) (JavbusCrawl, error) {
	collector := colly.NewCollector()
	collector.ParseHTTPErrorResponse = true
	collector.SetRedirectHandler(func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	})
	collector.Limit(&colly.LimitRule{
		Parallelism: 1,
		RandomDelay: 5 * time.Second,
	})
	collector.OnRequest(func(r *colly.Request) {
		logger.Infow("Visiting: " + r.URL.String())
	})
	http := resty.New()

	if config.Proxy != "" {
		_, err := url.Parse(config.Proxy)
		if err != nil {
			return nil, err
		}

		if err := collector.SetProxy(config.Proxy); err != nil {
			return nil, err
		}

		http.SetProxy(config.Proxy)
	}

	return &javbusCrawl{
		logger:     logger,
		opt:        opt,
		collector:  collector,
		httpClient: http,
	}, nil
}

// TODO: need query function
func (jc *javbusCrawl) GetJavbusMovieByHomePage() ([]JavMovie, error) {
	return nil, nil
}

func (jc *javbusCrawl) GetJavbusMovie() ([]JavMovie, error) {
	queue, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})

	for _, code := range jc.opt.Code {
		queue.AddURL(JAVBUS_URL + "/" + code)
	}

	jc.collector.OnHTML(".container", jc.javbusMovieInfoCrawl)
	jc.collector.OnHTML("body", jc.javbusMovieMagnetCrawl)
	queue.Run(jc.collector)
	return nil, nil
}

func (jc *javbusCrawl) GetJavbusMovieByPrefix() ([]JavMovie, error) {
	codes := jc.getCodeByPrefix()
	queue, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})
	for _, code := range codes {
		queue.AddURL(JAVBUS_URL + "/" + code)
	}

	jc.collector.OnHTML(".container", jc.javbusMovieInfoCrawl)
	jc.collector.OnHTML("body", jc.javbusMovieMagnetCrawl)
	queue.Run(jc.collector)
	return nil, nil
}

func (jc *javbusCrawl) GetJavbusMovieByStar() ([]JavMovie, error) {
	collector := jc.collector.Clone()
	pageCollector := jc.collector.Clone()
	queue, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})

	pageCollector.OnHTML("body", func(h *colly.HTMLElement) {
		h.ForEach("#waterfall > div", func(i int, element *colly.HTMLElement) {
			queue.AddURL(element.ChildAttr("a", "href"))
		})
		h.ForEach("div.text-center.hidden-xs > ul", func(i int, element *colly.HTMLElement) {
			page := element.ChildAttr("a#next", "href")
			if page != "" {
				pageCollector.Visit(JAVBUS_URL + element.ChildAttr("a#next", "href"))
			}
		})
	})

	for _, starCode := range jc.opt.StarCode {
		err := pageCollector.Visit(JAVBUS_URL + "/star/" + starCode)
		if err != nil {
			return nil, err
		}
	}
	pageCollector.Wait()

	if queue.IsEmpty() {
		return nil, nil
	}

	collector.OnHTML(".container", jc.javbusMovieInfoCrawl)
	collector.OnHTML("body", jc.javbusMovieMagnetCrawl)
	queue.Run(collector)
	collector.Wait()

	return nil, nil
}

func (jc *javbusCrawl) GetJavbusMovieByFilepath() ([]JavMovie, error) {
	var videoExt = []string{".avi", ".mp4", ".mkv"}
	var codes []string = make([]string, 0)
	if err := filepath.Walk(jc.opt.VideosPath, func(path string, fi os.FileInfo, err error) error {
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
				code := strings.Replace(filename, b, "", -1)
				codes = append(codes, code)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	fmt.Println(codes)

	return nil, nil
}

func (jc *javbusCrawl) SaveLocal(infos []JavMovie) error {
	df := dataframe.LoadStructs(infos)
	f, err := os.OpenFile(path.Join(jc.opt.OutPath, time.Now().Local().Format("2006-01-02-15-04-05")+"-jav_Infow.csv"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		jc.logger.Errorw("Save jav Infow file failed error: %s" + err.Error())
		return err
	}
	defer f.Close()
	if err := df.WriteCSV(f); err != nil {
		return err
	}

	// if len(jc.javMagnets) == 0 {
	// 	return nil
	// }

	// f, err := os.OpenFile(path.Join(jc.option.DestPath, time.Now().Local().Format("2006-01-02-15-04-05")+"-jav_magnet.text"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	// if err != nil {
	// 	jc.logger.Errorw("Save jav Infow file failed error: %s" + err.Error())
	// 	return err
	// }
	// defer f.Close()

	// for _, v := range jc.javMagnets {
	// 	f.WriteString(v + "\n")
	// }

	// var urlImg = ""
	// u, err := url.ParseRequestURI(coverPath)
	// if err != nil {
	// 	jc.logger.Errorw("parse cover path failed error: %s" + err.Error())
	// 	return err
	// }
	// if u.Host == "" {
	// 	urlImg = jc.javbusUrl + coverPath
	// } else {
	// 	urlImg = coverPath
	// }

	// jc.logger.Infow("downloading coverage url: " + urlImg)
	// ext := path.Ext(urlImg)
	// resp, err := jc.httpClient.Get(urlImg)
	// if err != nil {
	// 	jc.logger.Errorw("downloading coverage error: " + err.Error())
	// 	return err
	// }
	// body, _ := ioutil.ReadAll(resp.Body)

	// f, err := os.OpenFile(path.Join(jc.option.DestPath, "cover", code+ext), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()
	// if _, err := io.Copy(f, bytes.NewReader(body)); err != nil {
	// 	return err
	// }
	return nil
}

func (jc *javbusCrawl) javbusMovieInfoCrawl(e *colly.HTMLElement) {
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

	// jc. = append(t.javMovies, *Info)
}

func (jc *javbusCrawl) javbusMovieMagnetCrawl(e *colly.HTMLElement) {
	if !jc.opt.DownloadMagent {
		return
	}

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

		res, err := jc.httpClient.R().SetHeaders(map[string]string{
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
		// t.javMovies[0].Magnets = append(t.javMovies[0].Magnets, mList...)
	})
}

func (jc *javbusCrawl) getCodeByPrefix() []string {
	var codes []string = make([]string, 0)
	for _, code := range jc.opt.PrefixCode {
		for i := jc.opt.PrefixMinNo; i < jc.opt.PrefixMaxNo; i++ {
			strNum := strconv.FormatUint(i, 10)
			if len(strNum) >= int(jc.opt.PrefixZero) {
				codes = append(codes, code+strNum)
			} else {
				zerosStr := make([]byte, int(jc.opt.PrefixZero)-len(strNum))
				for i := range zerosStr {
					zerosStr[i] = '0'
				}
				codes = append(codes, code+string(append(zerosStr, []byte(strNum)...)))
			}
		}
	}
	return codes
}
