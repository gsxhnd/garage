package command

import (
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
	Action: func(context *cli.Context) error {
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
	}
	dbDirFlag = &cli.StringFlag{
		Name:    "dest",
		Aliases: nil,
		Usage:   "",
	}
)
