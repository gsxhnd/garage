package command

import (
	"fmt"
	"garage/core"
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
	Action: func(context *cli.Context) error {
		fmt.Println("sync")
		core.GetSubject()
		return nil
	},
}
