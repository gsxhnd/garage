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
var crawlCmd = &cli.Command{
	Name:        "crawl",
	Aliases:     nil,
	Usage:       "crawl jav data.",
	UsageText:   "crawl --site [javbus/javlibrary] -c XXX-001",
	Description: "crawl jav data, support javbus and javlibrary site.",
	ArgsUsage:   "",
	Flags: []cli.Flag{
		siteFlag,
	},
	Action: func(ctx *cli.Context) error {
		newLogger := utils.GetLogger()
		c := crawl.NewCrawlClient(newLogger)
		_ = c.SetProxy(ctx.String("proxy"))
		c.StarCrawlJavbusMovie(ctx.String("code"))
		return nil
	},
}
