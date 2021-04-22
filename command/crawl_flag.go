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
		Value:       "javbus",
	}
	baseFlag = &cli.StringFlag{
		Name:        "base",
		Aliases:     []string{"b"},
		Destination: nil,
		HasBeenSet:  false,
	}
)
