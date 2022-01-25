package command

import (
	"github.com/urfave/cli/v2"
)

var (
	proxyFlag = &cli.StringFlag{Name: "proxy", Usage: "代理配置"}
	siteFlag  = &cli.StringFlag{
		Name:        "site",
		Usage:       "选择爬取数据的网站",
		Destination: nil,
		HasBeenSet:  false,
		Value:       "javbus",
	}
)
