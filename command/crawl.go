package command

import (
	"fmt"
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
		searchFlag,
		siteFlag,
		baseFlag,
		starFlag,
	},
	Before: func(ctx *cli.Context) error {
		return nil
	},
	Action: func(ctx *cli.Context) error {
		fmt.Println("crawl")
		return nil
	},
}
