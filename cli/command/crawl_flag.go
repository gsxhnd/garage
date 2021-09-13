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
