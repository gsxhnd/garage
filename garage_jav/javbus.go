package garage_jav

import (
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

type JavbusCrawl interface {
	GetJavbusMovie(opt *JavbusOption) ([]JavMovie, error)           // 通过番号爬取对应的电影信息
	GetJavbusMovieByHomePage(opt *JavbusOption) ([]JavMovie, error) // 通过首页爬取对应的电影信息
	GetJavbusMovieByPrefix(opt *JavbusOption) ([]JavMovie, error)   // 通过番号前缀爬取对应的电影信息
	GetJavbusMovieByStar(opt *JavbusOption) ([]JavMovie, error)     // 通过演员ID爬取对应的电影信息
	GetJavbusMovieByFilepath(opt *JavbusOption) ([]JavMovie, error) // 访问文件夹下的视频列表爬取电影信息
	SaveLocal(destPath string) error
}

type javbusCrawl struct {
	logger        utils.Logger
	collector     *colly.Collector
	pageCollector *colly.Collector
	httpClient    *http.Client
	requestClient *resty.Client
	maxDepth      int
	javbusUrl     string
	javInfos      []JavMovie
	javMagnets    []string
}

func NewJavbusCrawl(logger utils.Logger, config *CrawlConfig) (JavbusCrawl, error) {
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

	httpClient := &http.Client{}

	if config.Proxy != "" {
		if err := collector.SetProxy(config.Proxy); err != nil {
			return nil, err
		}
		uri, err := url.Parse(config.Proxy)
		if err != nil {
			return nil, err
		}
		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(uri),
			},
		}
	}
	if err := utils.MakeDir(filepath.Join(config.DestPath, "cover")); err != nil {
		return nil, err
	}

	return &javbusCrawl{
		logger:        logger,
		collector:     collector,
		httpClient:    httpClient,
		requestClient: resty.New(),
		maxDepth:      100,
		javbusUrl:     "https://www.javbus.com",
		javInfos:      make([]JavMovie, 0),
		javMagnets:    make([]string, 0),
	}, nil
}

func (cc *javbusCrawl) GetJavbusMovieByHomePage(opt *JavbusOption) ([]JavMovie, error) {
	return cc.javInfos, nil
}

func (cc *javbusCrawl) GetJavbusMovie(opt *JavbusOption) ([]JavMovie, error) {
	collector := cc.collector.Clone()
	queue, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})

	if opt.DownloadMagent {
		collector.OnHTML("body", cc.getJavMovieMagnetByJavbus)
	}
	collector.OnHTML(".container", cc.getJavMovieInfowByJavbus)

	for _, v := range opt.Code {
		queue.AddURL(cc.javbusUrl + "/" + v)
	}

	err := queue.Run(cc.collector)
	if err != nil {
		return nil, err
	} else {
		return cc.javInfos, nil
	}
}

func (cc *javbusCrawl) GetJavbusMovieByPrefix(opt *JavbusOption) ([]JavMovie, error) {
	queue, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})
	collector := cc.collector.Clone()
	codes := cc.getCodeByPrefix(opt)

	for _, v := range codes {
		queue.AddURL(cc.javbusUrl + "/" + v)
	}

	if opt.DownloadMagent {
		collector.OnHTML("body", cc.getJavMovieMagnetByJavbus)
	}
	collector.OnHTML(".container", cc.getJavMovieInfowByJavbus)

	queue.Run(cc.collector)

	return cc.javInfos, nil
}

func (cc *javbusCrawl) GetJavbusMovieByStar(opt *JavbusOption) ([]JavMovie, error) {
	return cc.javInfos, nil
}

func (cc *javbusCrawl) GetJavbusMovieByFilepath(opt *JavbusOption) ([]JavMovie, error) {
	var videoExt = []string{".avi", ".mp4", ".mkv"}
	if err := filepath.Walk(opt.VideosPath, func(path string, fi os.FileInfo, err error) error {
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
				cc.collectorQueue.AddURL((cc.javbusUrl + "/" + filePrefix))
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	cc.collector.OnHTML(".container", cc.getJavMovieInfowByJavbus)
	if err := cc.collectorQueue.Run(cc.collector); err != nil {
		return nil, err
	}

	return cc.javInfos, nil
}

func (cc *javbusCrawl) SaveLocal(destPath string) error {
	df := dataframe.LoadStructs(cc.javInfos)
	f, err := os.OpenFile(path.Join(destPath, time.Now().Local().Format("2006-01-02-15-04-05")+"-jav_Infow.csv"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		cc.logger.Errorw("Save jav Infow file failed error: %s" + err.Error())
		return err
	}
	defer f.Close()
	return df.WriteCSV(f)
}

func (cc *javbusCrawl) getJavMovieInfowByJavbus(e *colly.HTMLElement) {
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
	cc.javInfos = append(cc.javInfos, *Infow)
}

func (cc *javbusCrawl) getJavMovieMagnetByJavbus(e *colly.HTMLElement) {
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
		cc.logger.Infow("Get magnet url: " + urlS)

		r, _ := http.NewRequest("GET", urlS, nil)
		r.Header.Add("Referer", e.Request.URL.Scheme+"://"+e.Request.URL.Host+e.Request.URL.Path)
		res, err := cc.httpClient.Do(r)
		if err != nil {
			cc.logger.Errorw("http client error: " + err.Error())
			return
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			cc.logger.Errorw("http read response error: " + err.Error())
			return
		}
		defer res.Body.Close()

		doc, err := htmlquery.Parse(strings.NewReader("<table><tbody>" + string(body) + "</tbody></table>"))
		if err != nil {
			cc.logger.Errorw("html query error: " + err.Error())
		}
		list, err := htmlquery.QueryAll(doc, "//tr")
		if err != nil {
			cc.logger.Errorw("html query tr error: " + err.Error())
		}
		if len(list) == 0 {
			cc.logger.Infow("当前无磁力连接")
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

func (cc *javbusCrawl) getJavStarMovieByJavbus(e *colly.HTMLElement) {
	queue, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})
	e.ForEach("#waterfall > div", func(i int, element *colly.HTMLElement) {
		queue.AddURL(element.ChildAttr("a", "href"))
	})
	e.ForEach("div.text-center.hidden-xs > ul", func(i int, element *colly.HTMLElement) {
		page := element.ChildAttr("a#next", "href")
		if page != "" {
			cc.pageCollector.Visit(cc.javbusUrl + element.ChildAttr("a#next", "href"))
		}
	})
}

func (cc *javbusCrawl) StartCrawlJavbusMovieByStar() error {
	starCode := cc.option.StarCode
	cc.logger.Debugw("Getting star code: " + starCode)
	cc.pageCollector = cc.collector.Clone()

	cc.pageCollector.OnHTML("body", cc.getJavStarMovieByJavbus)

	if err := cc.pageCollector.Visit(cc.javbusUrl + "/star/" + starCode); err != nil {
		return err
	}
	cc.pageCollector.Wait()
	if cc.collectorQueue.IsEmpty() {
		return nil
	}

	InfowCrawlClient := cc.collector.Clone()
	InfowCrawlClient.OnRequest(func(r *colly.Request) {
		cc.logger.Infow("Visiting: " + r.URL.String())
	})
	if cc.option.DownloadMagent {
		InfowCrawlClient.OnHTML("body", cc.getJavMovieMagnetByJavbus)
	}
	InfowCrawlClient.OnHTML(".container", cc.getJavMovieInfowByJavbus)
	cc.collectorQueue.Run(InfowCrawlClient)

	cc.collector.Wait()

	return nil
}

func (cc *javbusCrawl) saveMagents() error {
	// if len(cc.javMagnets) == 0 {
	// 	return nil
	// }

	// f, err := os.OpenFile(path.Join(cc.option.DestPath, time.Now().Local().Format("2006-01-02-15-04-05")+"-jav_magnet.text"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	// if err != nil {
	// 	cc.logger.Errorw("Save jav Infow file failed error: %s" + err.Error())
	// 	return err
	// }
	// defer f.Close()

	// for _, v := range cc.javMagnets {
	// 	f.WriteString(v + "\n")
	// }
	return nil
}

func (cc *javbusCrawl) saveCovers(coverPath, code string) error {
	// var urlImg = ""
	// u, err := url.ParseRequestURI(coverPath)
	// if err != nil {
	// 	cc.logger.Errorw("parse cover path failed error: %s" + err.Error())
	// 	return err
	// }
	// if u.Host == "" {
	// 	urlImg = cc.javbusUrl + coverPath
	// } else {
	// 	urlImg = coverPath
	// }

	// cc.logger.Infow("downloading coverage url: " + urlImg)
	// ext := path.Ext(urlImg)
	// resp, err := cc.httpClient.Get(urlImg)
	// if err != nil {
	// 	cc.logger.Errorw("downloading coverage error: " + err.Error())
	// 	return err
	// }
	// body, _ := ioutil.ReadAll(resp.Body)

	// f, err := os.OpenFile(path.Join(cc.option.DestPath, "cover", code+ext), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()
	// if _, err := io.Copy(f, bytes.NewReader(body)); err != nil {
	// 	return err
	// }
	return nil
}

func (cc *javbusCrawl) getCodeByPrefix(opt *JavbusOption) []string {
	var codes []string = make([]string, 0)
	for _, code := range codes {
		for i := opt.PrefixMinNo; i < opt.PrefixMaxNo; i++ {
			strNum := strconv.FormatUint(i, 10)
			if len(strNum) >= int(opt.PrefixZero) {
				codes = append(codes, code+strNum)
			} else {
				zerosStr := make([]byte, int(opt.PrefixZero)-len(strNum))
				for i := range zerosStr {
					zerosStr[i] = '0'
				}
				codes = append(codes, code+string(append(zerosStr, []byte(strNum)...)))
			}
		}
	}
	return codes
}
