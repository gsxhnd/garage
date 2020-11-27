package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"runtime"
)

var (
	gitTag       string
	gitCommit    string
	gitTreeState string
	buildDate    string
	goVersion    = runtime.Version()
	compiler     = runtime.Compiler
	platform     = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

var versionCmd = &cli.Command{
	Name:        "version",
	Usage:       "Show version",
	Description: "Show version",
	Action: func(context *cli.Context) error {
		fmt.Println("version: ", gitTag)
		fmt.Println("commit: ", gitCommit)
		fmt.Println("tree state: ", gitTreeState)
		fmt.Println("build date: ", buildDate)
		fmt.Println("go version: ", goVersion)
		fmt.Println("go compiler: ", compiler)
		fmt.Println("platform: ", platform)
		return nil
	},
}
