package main

import (
	"github.com/gsxhnd/garage/garage_di"
	"github.com/urfave/cli/v2"
)

var serverCmd = &cli.Command{
	Name:  "server",
	Usage: "",
	Action: func(context *cli.Context) error {
		app, _ := garage_di.InitApp()
		return app.Run()
	},
}
