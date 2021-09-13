package garage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func (c2 *Client) StarCrawlJavbusMovie(code, proxy string) {
	var data JavMovie
	var cover string
	c2.collector.OnHTML(".container", func(e *colly.HTMLElement) {
		data.Title = e.ChildText("h3")
		cover = e.ChildAttr(".screencap img", "src")
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
	})
	err := c2.collector.Visit(c2.javbusUrl + code)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	_, err = os.Stat("./javs/" + code)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir("./javs/"+code, os.ModePerm)
			if err != nil {
				return
			}
		} else {
			return
		}
	}

	if cover != "" {
		uri, _ := url.Parse(proxy)

		client := http.Client{
			Transport: &http.Transport{
				// 设置代理
				Proxy: http.ProxyURL(uri),
			},
		}
		resp, _ := client.Get(c2.javbusUrl + cover)
		body, _ := ioutil.ReadAll(resp.Body)
		out, _ := os.Create("./javs/" + code + "/" + code + ".jpg")
		io.Copy(out, bytes.NewReader(body))
	}
	saveData, _ := json.Marshal(&data)
	err = ioutil.WriteFile("./javs/"+code+"/info.json", saveData, os.ModeAppend)
	if err != nil {
		return
	}
}
