package crawl_cmd

import (
	"github.com/urfave/cli/v2"
)

var (
	proxyFlag  = &cli.StringFlag{Name: "proxy", Usage: "代理配置,如: http://127.0.0.1:1080"}
	outputFlag = &cli.StringFlag{Name: "output", Usage: "设置下载目录", Value: "./javbus"}
	siteFlag   = &cli.StringFlag{
		Name:        "site",
		Usage:       "选择爬取数据的网站,支持网站(javbus)",
		Destination: nil,
		HasBeenSet:  false,
		Value:       "javbus",
	}
	prefixCodeFlag = &cli.StringFlag{Name: "prefix-code", Required: true}
	prefixMinFlag  = &cli.IntFlag{Name: "prefix-min", Required: true, Value: 1}
	prefixMaxFlag  = &cli.IntFlag{Name: "prefix-max", Required: true, Value: 5}
)
