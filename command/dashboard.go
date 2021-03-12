package command

import (
	"garage/api"
	"garage/dao"
	"github.com/gsxhnd/owl"
	"github.com/urfave/cli/v2"
	"time"
)

// start dashboard api
var dashboardCmd = &cli.Command{
	Name:        "dashboard",
	Aliases:     nil,
	Usage:       "dashboard",
	UsageText:   "dashboard",
	Description: "start web ui",
	Before: func(ctx *cli.Context) error {
		conf := ctx.String("conf")
		owl.SetConfName(conf)
		err := owl.ReadConf()
		if err != nil {
			return err
		} else {
			return nil
		}
	},
	Action: func(ctx *cli.Context) error {
		var (
			port   = owl.GetString("dashboard.port")
			imgDir = owl.GetString("img_dir")
		)
		if port == "" {
			port = ":8080"
		}
		if imgDir == "" {
			imgDir = "./img"
		}

		defer dao.Database.Close()
		switch owl.GetString("db.source") {
		case "sqlite":
			err := dao.Database.ConnectSQLite(owl.GetString("db.sqlite.file"))
			if err != nil {
				return err
			}
		}

		if ctx.Bool("open") {
			go time.AfterFunc(2*time.Second, func() {
				api.OpenBrowser("http://localhost:" + port)
			})
		}

		err := api.Run(port, imgDir)
		return err
	},
	OnUsageError: nil,
	Flags: []cli.Flag{
		confFlag,
		openFlag,
	},
}
