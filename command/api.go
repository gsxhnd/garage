package command

import (
	"errors"
	"garage/api"
	"garage/dao"
	"garage/model"
	"github.com/gsxhnd/owl"
	"github.com/urfave/cli/v2"
	"log"
)

// start dashboard api
var apiCmd = &cli.Command{
	Name:        "api",
	Aliases:     nil,
	Usage:       "api",
	UsageText:   "api",
	Description: "start api server",
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
		case "postgre":
			err := dao.Database.ConnectPostgreSQL()
			if err != nil {
				return err
			}
		case "mysql":
			err := dao.Database.ConnectMySQL()
			if err != nil {
				return err
			}
		default:
			return errors.New("err database sourse")
		}

		err := dao.Database.Default.AutoMigrate(
			&model.JavMovie{},
			&model.JavStar{},
			&model.JavTag{},
			&model.JavMovieSatr{},
			&model.JavMovieTag{},
		)
		if err != nil {
			log.Println(err)
		}

		err = api.Run(port, imgDir)
		return err
	},
	OnUsageError: nil,
	Flags: []cli.Flag{
		confFlag,
	},
}
