package command

import "github.com/urfave/cli/v2"

var (
	dbDirFlag = &cli.StringFlag{
		Name:    "db-dir",
		Aliases: nil,
		Usage:   "",
	}
	dbNameFlag = &cli.StringFlag{
		Name:        "db-name",
		Usage:       "",
		EnvVars:     nil,
		FilePath:    "",
		Value:       "jav.db",
		DefaultText: "",
	}
)
