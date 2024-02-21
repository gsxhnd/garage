package main

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

var (
	tag       string
	buildDate string
	gitCommit string
)

var versionCmd = &cli.Command{
	Name:  "version",
	Usage: "显示版本号",
	Action: func(context *cli.Context) error {
		fmt.Println("Version: ", tag)
		fmt.Println("BuildDate:", buildDate)
		fmt.Println("GitCommit:", gitCommit)
		fmt.Println("GoVersion:", runtime.Version())
		fmt.Println("Compiler:", runtime.Compiler)
		fmt.Println("Paltform:", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
		return nil
	},
}
