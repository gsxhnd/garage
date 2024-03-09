package garage_jav

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-resty/resty/v2"
	"github.com/gocolly/colly/v2"
	"github.com/gsxhnd/garage/utils"
)

type JavCrawl interface {
	GetJavbusMovie(opt *JavbusOption) ([]JavMovie, error)           // 通过番号爬取对应的电影信息
	GetJavbusMovieByHomePage(opt *JavbusOption) ([]JavMovie, error) // 通过首页爬取对应的电影信息
	GetJavbusMovieByPrefix(opt *JavbusOption) ([]JavMovie, error)   // 通过番号前缀爬取对应的电影信息
	GetJavbusMovieByStar(opt *JavbusOption) ([]JavMovie, error)     // 通过演员ID爬取对应的电影信息
	GetJavbusMovieByFilepath(opt *JavbusOption) ([]JavMovie, error) // 访问文件夹下的视频列表爬取电影信息
	SaveLocal(destPath string, infos []JavMovie) error
}

type javCrawl struct {
	logger     utils.Logger
	collector  *colly.Collector
	httpClient *resty.Client
}

func NewJavCrawl(logger utils.Logger, config *CrawlConfig) (JavCrawl, error) {
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

	return &javCrawl{
		logger:     logger,
		collector:  collector,
		httpClient: http,
	}, nil
}

func (cc *javCrawl) GetJavbusMovieByHomePage(opt *JavbusOption) ([]JavMovie, error) {
	return nil, nil
}

// collector.OnHTML("body", cc.getJavMovieMagnetByJavbus)
func (cc *javCrawl) GetJavbusMovie(opt *JavbusOption) ([]JavMovie, error) {
	return nil, nil
}

func (cc *javCrawl) GetJavbusMovieByPrefix(opt *JavbusOption) ([]JavMovie, error) {
	// codes := cc.getCodeByPrefix(opt)

	return nil, nil
}

func (cc *javCrawl) GetJavbusMovieByStar(opt *JavbusOption) ([]JavMovie, error) {
	return nil, nil
}

func (cc *javCrawl) GetJavbusMovieByFilepath(opt *JavbusOption) ([]JavMovie, error) {
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
				code := strings.Replace(filename, b, "", -1)
				opt.Code = append(opt.Code, code)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	fmt.Println(opt.Code)

	return nil, nil
}

func (cc *javCrawl) SaveLocal(destPath string, infos []JavMovie) error {
	df := dataframe.LoadStructs(infos)
	f, err := os.OpenFile(path.Join(destPath, time.Now().Local().Format("2006-01-02-15-04-05")+"-jav_Infow.csv"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		cc.logger.Errorw("Save jav Infow file failed error: %s" + err.Error())
		return err
	}
	defer f.Close()
	return df.WriteCSV(f)
}

func (cc *javCrawl) StartCrawlJavbusMovieByStar() error {
	// starCode := cc.option.StarCode
	// cc.logger.Debugw("Getting star code: " + starCode)
	// cc.pageCollector = cc.collector.Clone()

	// cc.pageCollector.OnHTML("body", cc.getJavStarMovieByJavbus)

	// if err := cc.pageCollector.Visit(cc.javbusUrl + "/star/" + starCode); err != nil {
	// 	return err
	// }
	// cc.pageCollector.Wait()
	// if cc.collectorQueue.IsEmpty() {
	// 	return nil
	// }

	// InfowCrawlClient := cc.collector.Clone()
	// InfowCrawlClient.OnRequest(func(r *colly.Request) {
	// 	cc.logger.Infow("Visiting: " + r.URL.String())
	// })
	// if cc.option.DownloadMagent {
	// 	InfowCrawlClient.OnHTML("body", cc.getJavMovieMagnetByJavbus)
	// }
	// InfowCrawlClient.OnHTML(".container", cc.getJavMovieInfowByJavbus)
	// cc.collectorQueue.Run(InfowCrawlClient)

	// cc.collector.Wait()

	return nil
}

func (cc *javCrawl) saveMagents() error {
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

func (cc *javCrawl) saveCovers(coverPath, code string) error {
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

func (cc *javCrawl) getCodeByPrefix(opt *JavbusOption) []string {
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
