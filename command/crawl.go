package command

import (
	"github.com/gsxhnd/owl"
	"github.com/urfave/cli/v2"
)

// crawl data
var crawlCmd = &cli.Command{
	Name:        "crawl",
	Aliases:     nil,
	Usage:       "crawl jav data.",
	UsageText:   "crawl --site [javbus/javlibrary] -s XXX-001",
	Description: "crawl jav data, support javbus and javlibrary site.",
	ArgsUsage:   "",
	Flags: []cli.Flag{
		confFlag,
		searchFlag,
		siteFlag,
		baseFlag,
	},
	Before: func(ctx *cli.Context) error {
		conf := ctx.String("conf")
		owl.SetConfName(conf)
		err := owl.ReadConf()
		if err != nil {
			return err
		} else {
			return nil
		}
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
