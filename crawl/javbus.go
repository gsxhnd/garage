package crawl

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

func (c2 *Client) StarCrawlJavbusMovie(code string) {
	info, err := c2.DownloadInfo(code)
	if err != nil {
		return
	}
	err = c2.DownloadCover(info.Code, info.Cover)
	if err != nil {
		return
	}
}

func (c2 *Client) DownloadInfo(code string) (*JavMovie, error) {
	c2.logger.Info("Download info: " + code)
	var data JavMovie
	c2.collector.OnHTML(".container", func(e *colly.HTMLElement) {
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
			//href := element.Attr("href")
			star := element.Attr("title")
			data.Stars = append(data.Stars, star)
		})

		// e.ForEach("#magnet-table , tr", func(i int, h *colly.HTMLElement) {
		// 	fmt.Println(i)
		// 	fmt.Println(h)
		// })
	})
	err := c2.collector.Visit(c2.javbusUrl + code)
	if err != nil {
		return nil, err
	}
	saveData, _ := json.Marshal(&data)
	err = ioutil.WriteFile("./javs/"+code+"/info.json", saveData, os.ModeAppend)
	if err != nil {
		c2.logger.Error("", zap.Error(err))
		return nil, err
	} else {
		return &data, nil
	}
}

func (c2 *Client) DownloadCover(code, cover string) error {
	resp, _ := c2.httpClient.Get(c2.javbusUrl + cover)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create("./javs/" + code + "/" + code + ".jpg")
	io.Copy(out, bytes.NewReader(body))
	return nil
}
