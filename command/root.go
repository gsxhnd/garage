package command

import (
	"github.com/urfave/cli/v2"
)

var (
	RootCmd = cli.NewApp()
)

func init() {
	RootCmd.HideVersion = true
	RootCmd.Usage = "JAV命令行工具"
	RootCmd.Flags = []cli.Flag{
		proxyFlag,
	}
	RootCmd.Commands = []*cli.Command{
		codeCmd,
		starCmd,
		prefixCmd,
		versionCmd,
	}
}
