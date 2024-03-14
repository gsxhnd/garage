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
		if app, err := garage_di.InitApp(ctx.Path("config")); err != nil {
			return err
		} else {
			return app.Run()
		}
	},
}
