package command

import (
	"github.com/urfave/cli/v2"
)

var (
	RootCmd = cli.NewApp()
)

func init() {
	RootCmd.Usage = "garage"
	RootCmd.Version = ""
	RootCmd.HideVersion = true
	RootCmd.Flags = []cli.Flag{
		proxyFlag,
	}
	RootCmd.Commands = []*cli.Command{
		crawlCmd,
		versionCmd,
	}
}
