package main

import (
	"os"

	"github.com/gsxhnd/garage/utils"
	"github.com/urfave/cli/v2"
)

var (
	rootCmd = cli.NewApp()
	logger  = utils.NewLogger(nil)
)

func init() {
	rootCmd.HideVersion = true
	rootCmd.Usage = "Set of crwal tool"
	rootCmd.Flags = []cli.Flag{}
	rootCmd.Commands = []*cli.Command{
		crawlJavbusCmd,
		crawlJavDBCmd,
		versionCmd,
	}
}

func main() {
	err := rootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
