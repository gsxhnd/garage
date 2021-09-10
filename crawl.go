package main

import (
	"github.com/gocolly/colly/v2"
)

type Client struct {
	collector *colly.Collector
	proxy     string
	maxDepth  int
	javbusUrl string
	javlibUrl string
}

type JavMovie struct {
	Code           string `json:"code"`
	Title          string `json:"title"`
	ProduceCompany string `json:"produce_company"`
	PublishCompany string `json:"publish_company"`
}

func NewCrawlClient() *Client {
	return &Client{
		collector: colly.NewCollector(),
		proxy:     "",
		maxDepth:  100,
		javbusUrl: "https://www.javbus.com/",
		javlibUrl: "https://www.javbus.com/",
	}
}

func (c2 *Client) SetProxy(proxy string) (err error) {
	err = c2.collector.SetProxy(proxy)
	return
}
