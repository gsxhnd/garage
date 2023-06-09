package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var versionCmd = &cli.Command{
	Name:  "version",
	Usage: "显示版本号",
	Action: func(context *cli.Context) error {
		fmt.Println("version cmd")
		return nil
	},
}
