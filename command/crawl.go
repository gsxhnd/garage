package command

import (
	"fmt"
	"garage/dao"
	"github.com/gocolly/colly/v2"
	"github.com/urfave/cli/v2"
)

// crawl data
var crawlCmd = &cli.Command{
	Name:         "crawl",
	Aliases:      nil,
	Usage:        "crawl jav data.",
	UsageText:    "crawl --site [javbus/javlibrary] -s XXX-001",
	Description:  "crawl  jav data, support javbus and javlibrary site.",
	ArgsUsage:    "",
	Category:     "",
	BashComplete: nil,
	Flags: []cli.Flag{
		searchFlag,
		siteFlag,
		baseFlag,
		proxyFlag,
		syncFlag,
		dbDirFlag,
	},
	Before: func(ctx *cli.Context) error {
		sync := ctx.Bool("sync")
		fmt.Println(sync)
		if sync {
			_ = dao.Database.Connect()
			defer dao.Database.Close()
		}
		return nil
	},
	Action: func(ctx *cli.Context) error {
		c := colly.NewCollector()
		_ = c.SetProxy(ctx.String("proxy"))
		c.MaxDepth = 100
		c.OnHTML(".container-fluid", func(e *colly.HTMLElement) {
			fmt.Println(e)
		})
		_ = c.Visit("https://www.javbus.com")
		return nil
	},
}