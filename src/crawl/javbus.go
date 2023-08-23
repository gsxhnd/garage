package crawl

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/inhies/go-bytesize"
	"go.uber.org/zap"
)

type JavbusCrawl interface { // 通过番号爬取对应的电影信息
	StartCrawlJavbusMovie(code string) error
	// 通过番号前缀爬取对应的电影信息
	StartCrawlJavbusMovieByPrefix() error
	// 通过演员ID爬取对应的电影信息
	StartCrawlJavbusMovieByStar(starCode string) error
	// 访问文件夹下的视频列表爬取电影信息
	StartCrawlJavbusMovieByFilepath(inputPath string) error
	// 设置代理
	setProxy(proxy string) error
	// 创建输出目录
	mkAllDir() error
	getJavMovieInfoByJavbus(element *colly.HTMLElement)
	getJavMovieMagnetByJavbus(e *colly.HTMLElement)
	getJavStarMovieByJavbus(e *colly.HTMLElement)

	// 保存CSV格式的电影信息
	saveJavInfos() error
	// 下载电影封面
	saveCovers(coverPath, name string) error
	// 下载磁力列表
	saveMagents() error
}

type javbusCrawl struct {
	logger         *zap.Logger
	collector      *colly.Collector
	pageCollector  *colly.Collector
	httpClient     *http.Client
	maxDepth       int
	javbusUrl      string
	javInfos       []JavMovie
	javMagnets     []string
	downloadMagent bool
	destPath       string
	starCode       string
	starPage       int
	prefixCode     string
	prefixMinNo    int
	prefixMaxNo    int
	javQueue       *queue.Queue
}

type CrawlOptions struct {
	DestPath       string
	Proxy          string
	DownloadMagent bool
	StarCode       string
	PrefixCode     string
	PrefixMinNo    int
	PrefixMaxNo    int
}

func NewJavbusCrawl(logger *zap.Logger, option CrawlOptions) (JavbusCrawl, error) {
	q, _ := queue.New(1, &queue.InMemoryQueueStorage{MaxSize: 10000})
	var client = &javbusCrawl{
		collector:      colly.NewCollector(),
		pageCollector:  nil,
		httpClient:     &http.Client{},
		maxDepth:       100,
		javbusUrl:      "https://www.javbus.com",
		logger:         logger,
		javInfos:       make([]JavMovie, 0),
		javMagnets:     make([]string, 0),
		downloadMagent: option.DownloadMagent,
		destPath:       option.DestPath,
		prefixCode:     option.PrefixCode,
		prefixMinNo:    option.PrefixMinNo,
		prefixMaxNo:    option.PrefixMaxNo,
		starPage:       2,
		javQueue:       q,
	}

	client.collector.Limit(&colly.LimitRule{
		Parallelism: 1,
		RandomDelay: 5 * time.Second,
	})
	if option.Proxy != "" {
		if err := client.setProxy(option.Proxy); err != nil {
			return nil, err
		}
	}

	if err := client.mkAllDir(); err != nil {
		return nil, err
	} else {
		return client, nil
	}
}

func (cc *javbusCrawl) setProxy(proxy string) error {
	if err := cc.collector.SetProxy(proxy); err != nil {
		return err
	}
	uri, err := url.Parse(proxy)
	if err != nil {
		return err
	}
	cc.httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},
	}
	return nil
}

func (cc *javbusCrawl) mkAllDir() error {
	fullPath := filepath.Join(cc.destPath, "cover")
	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(fullPath, os.ModePerm)
			if err != nil {
				return err
			}
		} else if os.IsExist(err) {
			return nil
		} else {
			return err
		}
	}
	return nil
}

func (cc *javbusCrawl) getJavMovieInfoByJavbus(e *colly.HTMLElement) {
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

func (cc *javbusCrawl) getJavStarMovieByJavbus(e *colly.HTMLElement) {
	e.ForEach("#waterfall > div", func(i int, element *colly.HTMLElement) {
		cc.javQueue.AddURL(element.ChildAttr("a", "href"))
	})
	e.ForEach("div.text-center.hidden-xs > ul", func(i int, element *colly.HTMLElement) {
		page := element.ChildAttr("a#next", "href")
		if page != "" {
			cc.pageCollector.Visit(cc.javbusUrl + element.ChildAttr("a#next", "href"))
		}
	})
}

func (cc *javbusCrawl) StartCrawlJavbusMovie(code string) error {
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

func (cc *javbusCrawl) StartCrawlJavbusMovieByPrefix() error {
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
	cc.saveMagents()
	for _, v := range cc.javInfos {
		err := cc.saveCovers(v.Cover, v.Code)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cc *javbusCrawl) StartCrawlJavbusMovieByStar(starCode string) error {
	cc.starCode = starCode
	cc.logger.Debug("Getting star code: " + starCode)
	cc.pageCollector = cc.collector.Clone()

	cc.pageCollector.OnRequest(func(r *colly.Request) {
		cc.logger.Info("Visiting: " + r.URL.String())
	})
	cc.pageCollector.OnHTML("body", cc.getJavStarMovieByJavbus)

	if err := cc.pageCollector.Visit(cc.javbusUrl + "/star/" + starCode); err != nil {
		return err
	}
	cc.pageCollector.Wait()
	if cc.javQueue.IsEmpty() {
		return nil
	}

	infoCrawlClient := cc.collector.Clone()
	infoCrawlClient.OnRequest(func(r *colly.Request) {
		cc.logger.Info("Visiting: " + r.URL.String())
	})
	if cc.downloadMagent {
		infoCrawlClient.OnHTML("body", cc.getJavMovieMagnetByJavbus)
	}
	infoCrawlClient.OnHTML(".container", cc.getJavMovieInfoByJavbus)
	cc.javQueue.Run(infoCrawlClient)

	cc.collector.Wait()

	cc.saveJavInfos()
	cc.saveMagents()
	for _, v := range cc.javInfos {
		err := cc.saveCovers(v.Cover, v.Code)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cc *javbusCrawl) StartCrawlJavbusMovieByFilepath(inputPath string) error {
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
