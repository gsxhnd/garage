package command

import (
	"github.com/gsxhnd/garage"
	"github.com/urfave/cli/v2"
	"os"
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
	Before: func(ctx *cli.Context) error {
		Logger.Debug("checking javs dir exists...")
		_, err := os.Stat("./javs")
		if err != nil {
			if os.IsNotExist(err) {
				err := os.Mkdir("./javs", os.ModePerm)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
		return nil
	},
	Action: func(ctx *cli.Context) error {
		c := garage.NewCrawlClient()
		_ = c.SetProxy(ctx.String("proxy"))
		c.StarCrawlJavbusMovie(ctx.String("code"), ctx.String("proxy"))
		return nil
	},
}
