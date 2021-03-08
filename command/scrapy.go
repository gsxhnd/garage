package command

import (
	"fmt"
	"garage/dao"
	"github.com/gocolly/colly/v2"
	"github.com/urfave/cli/v2"
)

//
var crawlCmd = &cli.Command{
	Name:         "crawl",
	Aliases:      nil,
	Usage:        "",
	UsageText:    "",
	Description:  "",
	ArgsUsage:    "",
	Category:     "",
	BashComplete: nil,
	Flags: []cli.Flag{
		searchFlag,
		siteFlag,
		baseFlag,
		proxyFlag,
		syncDirFlag,
		dbDirFlag,
	},
	Before: func(context *cli.Context) error {
		_ = dao.Database.Connect()
		defer dao.Database.Close()
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
		Destination: nil,
		HasBeenSet:  false,
	}
	baseFlag = &cli.StringFlag{
		Name:        "base",
		Aliases:     []string{"b"},
		Destination: nil,
		HasBeenSet:  false,
	}
	proxyFlag = &cli.StringFlag{
		Name:        "proxy",
		Destination: nil,
		HasBeenSet:  false,
	}
	syncDirFlag = &cli.BoolFlag{
		Name:  "sync",
		Usage: "",
		Value: false,
	}
	dbDirFlag = &cli.StringFlag{
		Name:    "dest",
		Aliases: nil,
		Usage:   "",
	}
)
