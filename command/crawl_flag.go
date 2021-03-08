package command

import "github.com/urfave/cli/v2"

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
	syncFlag = &cli.BoolFlag{
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
