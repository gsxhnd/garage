package command

import (
	"github.com/gsxhnd/garage/crawl"
	"github.com/gsxhnd/garage/utils"
	"github.com/urfave/cli/v2"
)

var (
	searchFlag = &cli.StringFlag{
		Name:        "search",
		Aliases:     []string{"s"},
		Usage:       "",
		Destination: nil,
		HasBeenSet:  false,
	}
	siteFlag = &cli.StringFlag{
		Name:        "site",
		Usage:       "选择爬取数据的网站",
		Destination: nil,
		HasBeenSet:  false,
		Value:       "javbus",
	}
	codeFlag = &cli.StringFlag{
		Name:        "code",
		Aliases:     []string{"c"},
		Usage:       "-c xxx-001",
		Destination: nil,
		HasBeenSet:  false,
	}
	starFlag = &cli.StringFlag{
		Name:        "star",
		Aliases:     []string{"t"},
		Destination: nil,
		HasBeenSet:  false,
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
		searchFlag,
		siteFlag,
		codeFlag,
		starFlag,
	},
	Action: func(ctx *cli.Context) error {
		newLogger := utils.GetLogger()
		c := crawl.NewCrawlClient(newLogger)
		_ = c.SetProxy(ctx.String("proxy"))
		c.StarCrawlJavbusMovie(ctx.String("code"))
		return nil
	},
}
