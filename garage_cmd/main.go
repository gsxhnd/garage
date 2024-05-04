package main

import (
	"os"

	"github.com/gsxhnd/garage/utils"
	"github.com/urfave/cli/v2"

	_ "github.com/glebarez/go-sqlite"
)

var (
	RootCmd = cli.NewApp()
	logger  = utils.NewLogger(&utils.Config{
		Dev: true,
		LogConfig: utils.LogConfig{
			Level: "debug",
		},
	})
)

func init() {
	RootCmd.HideVersion = true
	RootCmd.Usage = "Set of crwal tool"
	RootCmd.Flags = []cli.Flag{}
	RootCmd.Commands = []*cli.Command{
		crawlJavbusCmd,
		versionCmd,
		serverCmd,
	}
}

func main() {
	err := RootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
