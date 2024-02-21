package main

import (
	"github.com/gsxhnd/garage/di"
	"github.com/urfave/cli/v2"
)

var serverCmd = &cli.Command{
	Name:  "server",
	Usage: "",
	Action: func(context *cli.Context) error {
		app, _ := di.InitApp()
		return app.Run()
	},
}
