package command

import (
	"github.com/gsxhnd/garage/crawl"
	"github.com/gsxhnd/garage/utils"
	"github.com/urfave/cli/v2"
)

var (
	siteFlag = &cli.StringFlag{
		Name:        "site",
		Usage:       "选择爬取数据的网站",
		Destination: nil,
		HasBeenSet:  false,
		Value:       "javbus",
	}
)

// crawl data
var codeCmd = &cli.Command{
	Name:        "code",
	Aliases:     nil,
	Usage:       "根据指定番号爬取数据",
	UsageText:   "crawl --site [javbus/javlibrary] XXX-001",
	Description: "crawl jav data, support javbus and javlibrary site.",
	ArgsUsage:   "",
	Flags: []cli.Flag{
		siteFlag,
	},
	Action: func(ctx *cli.Context) error {
		newLogger := utils.GetLogger()
		var code = ctx.Args().Get(0)
		c := crawl.NewCrawlClient(newLogger)
		_ = c.SetProxy(ctx.String("proxy"))
		c.StarCrawlJavbusMovie(code)
		return nil
	},
}
