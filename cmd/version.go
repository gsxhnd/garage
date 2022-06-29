package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var versionCmd = &cli.Command{
	Name: "version",
	Action: func(context *cli.Context) error {
		fmt.Println("version cmd")
		return nil
	},
}
