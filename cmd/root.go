package cmd

import (
	"github.com/urfave/cli/v2"
)

var (
	RootCmd = cli.NewApp()
)

func init() {
	RootCmd.HideVersion = true
	RootCmd.Usage = "命令行工具"
	RootCmd.Flags = []cli.Flag{}
	RootCmd.Commands = []*cli.Command{
		javCodeCmd,
		javStarCmd,
		javPrefixCmd,
		videoConvertCmd,
		videoSubtitleCmd,
		videoExtractSubtitleCmd,
		versionCmd,
	}
}
