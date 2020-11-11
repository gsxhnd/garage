package cmd

import "github.com/urfave/cli/v2"

var (
	App = cli.NewApp()
)

func init() {
	App.Version = "1"
}
