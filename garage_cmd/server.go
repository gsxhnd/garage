package main

import (
	"github.com/gsxhnd/garage/garage_di"
	"github.com/urfave/cli/v2"
)

var serverCmd = &cli.Command{
	Name:  "server",
	Usage: "",
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name: "config",
		},
	},
	Action: func(ctx *cli.Context) error {
		app, _ := garage_di.InitApp(ctx.Path("config"))
		return app.Run()
	},
}
