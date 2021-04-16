package command

import (
	"fmt"
	"garage/utils"
	"github.com/urfave/cli/v2"
)

var versionCmd = &cli.Command{
	Name:        "version",
	Usage:       "Show version",
	Description: "Show version",
	Action: func(context *cli.Context) error {
		var info = utils.GetVersionInfo()
		fmt.Println("version:     ", info.Version)
		fmt.Println("commit:      ", info.GitCommit)
		fmt.Println("tree state:  ", info.GitTreeState)
		fmt.Println("build date:  ", info.BuildDate)
		fmt.Println("go version:  ", info.GoVersion)
		fmt.Println("go compiler: ", info.Compiler)
		fmt.Println("platform:    ", info.Platform)
		return nil
	},
}
