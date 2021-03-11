package command

import (
	"garage/crawl"
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
	},
	Before: func(ctx *cli.Context) error {
		var (
			proxy = ctx.String("proxy")
		)
		crawl.SetProxy(proxy)
		return nil
	},
	Action: func(ctx *cli.Context) error {
		crawl.Tt()
		return nil
	},
}
