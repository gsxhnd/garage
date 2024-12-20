package cmd

import "github.com/urfave/cli/v2"

var CrawlJavbusCmd = &cli.Command{
	Name:        "crawl_javbus",
	Description: "",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "proxy", Usage: "代理配置,如: http://127.0.0.1:1080"},
		&cli.StringFlag{Name: "dest", Usage: "设置下载目录", Value: "./javbus"},
		&cli.BoolFlag{
			Name:  "magnet",
			Usage: "保存磁力链接,开启参数 --magnet",
			Value: false,
		},
	},
}
