package command

import (
	"fmt"
	"garage/core"
	"garage/dao"
	"github.com/urfave/cli/v2"
)

var syncCmd = &cli.Command{
	Name:         "sync",
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
	},
	Action: func(context *cli.Context) error {
		fmt.Println("sync")
		err := dao.Database.Connect()
		if err != nil {
			return err
		}
		dao.GetSubject()

		core.GetSubject()
		return nil
	},
}

var searchFlag = &cli.StringFlag{
	Name:        "search",
	Aliases:     []string{"s"},
	Destination: nil,
	HasBeenSet:  false,
}

var baseFlag = &cli.StringFlag{
	Name:        "base",
	Aliases:     []string{"b"},
	Destination: nil,
	HasBeenSet:  false,
}

var proxyFlag = &cli.StringFlag{
	Name:        "proxy",
	Destination: nil,
	HasBeenSet:  false,
}
