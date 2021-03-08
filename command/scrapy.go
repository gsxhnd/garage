package command

import (
	"garage/dao"
	"github.com/urfave/cli/v2"
)

var scrapyCmd = &cli.Command{
	Name:         "scrapy",
	Aliases:      nil,
	Usage:        "",
	UsageText:    "",
	Description:  "",
	ArgsUsage:    "",
	Category:     "",
	BashComplete: nil,
	Before:       nil,
	After:        nil,
	Flags: []cli.Flag{
		searchFlag,
		baseFlag,
		proxyFlag,
		syncDirFlag,
		dbDirFlag,
	},
	Action: func(ctx *cli.Context) error {
		_ = dao.Database.Connect()
		defer dao.Database.Close()
		dao.GetJavMovie()
		//fmt.Println(ctx.String("search"))
		//fmt.Println(ctx.Bool("sync"))
		//c := colly.NewCollector()
		//c.MaxDepth = 100
		//c.OnHTML(".pin-list-wrap", func(e *colly.HTMLElement) {
		//	fmt.Println(e)
		//})
		//_ = c.Visit("https://juejin.cn/pins/recommended")
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
