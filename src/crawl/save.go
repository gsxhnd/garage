package crawl

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

func (cc *crawlClient) DownloadInfo(code string) (*JavMovie, error) {
	cc.logger.Info("Download info: " + code)
	var data JavMovie
	cc.collector.OnHTML(".container", func(e *colly.HTMLElement) {
		data.Title = e.ChildText("h3")
		data.Cover = e.ChildAttr(".screencap img", "src")
		e.ForEach(".info p", func(i int, element *colly.HTMLElement) {
			key := element.ChildText("span")
			switch i {
			case 0:
				data.Code = element.ChildTexts("span")[1]
			}
			switch key {
			case "發行日期:":
				pd := element.Text
				data.PublishDate = strings.Split(pd, " ")[1]
			case "長度:":
				pd := element.Text
				p := strings.Split(pd, " ")[1]
				data.Length = strings.Split(p, "分鐘")[0]
			case "導演:":
				data.Director = element.ChildText("a")
			case "製作商:":
				data.ProduceCompany = element.ChildText("a")
			case "發行商:":
				data.PublishCompany = element.ChildText("a")
			case "系列:":
				data.Series = element.ChildText("a")
			}
		})
		e.ForEach("ul li .star-name a", func(i int, element *colly.HTMLElement) {
			// star := element.Attr("title")
			// data.Stars = append(data.Stars, star)
		})
	})
	err := cc.collector.Visit(cc.javbusUrl + code)
	if err != nil {
		return nil, err
	}
	saveData, _ := json.Marshal(&data)
	err = ioutil.WriteFile("./javs/"+code+"/info.json", saveData, os.ModeAppend)
	if err != nil {
		cc.logger.Error("", zap.Error(err))
		return nil, err
	} else {
		return &data, nil
	}
}

func (cc *crawlClient) DownloadCover(code, cover string) error {
	resp, _ := cc.httpClient.Get(cc.javbusUrl + cover)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create("./javs/" + code + "/" + code + ".jpg")
	io.Copy(out, bytes.NewReader(body))
	return nil
}

func (cc *crawlClient) saveJavInfos() error {
	var data = []JavMovie{
		{
			Code: "test-01",
		},
	}
	df := dataframe.LoadStructs(data)
	f, err := os.OpenFile("./../../test", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		// cc.logger.Error("")
		return err
	}
	return df.WriteCSV(f)
}

func (cc *crawlClient) saveCovers() error {
	return nil
}
func (cc *crawlClient) saveMagents() error {
	return nil
}
