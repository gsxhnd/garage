package main

import (
	"os"

	"github.com/gsxhnd/garage/utils"
	"github.com/urfave/cli/v2"
)

var (
	RootCmd = cli.NewApp()
	logger  = utils.NewLogger()
)

func init() {
	RootCmd.HideVersion = true
	RootCmd.Usage = "Set of crwal tool"
	RootCmd.Flags = []cli.Flag{}
	RootCmd.Commands = []*cli.Command{
		crawlJavbusCmd,
		crawlJavDBCmd,
		crawlTenhouCmd,
		versionCmd,
	}
}

func main() {
	err := RootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
