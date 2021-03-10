package command

import (
	"garage/api"
	"github.com/urfave/cli/v2"
	"time"
)

// start dashboard api
var dashboardCmd = &cli.Command{
	Name:         "dashboard",
	Aliases:      nil,
	Usage:        "dashboard",
	UsageText:    "dashboard",
	Description:  "start web ui",
	ArgsUsage:    "",
	Category:     "",
	BashComplete: nil,
	Before: func(ctx *cli.Context) error {
		go time.AfterFunc(2*time.Second, func() {
			api.OpenBrowser("http://localhost:" + ctx.String("port"))
		})
		return nil
	},
	After: func(ctx *cli.Context) error {
		api.Run(ctx.String("port"))
		return nil
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
	OnUsageError: nil,
	Flags: []cli.Flag{
		portFlag,
		coverImgFlag,
		starImgFlag,
		dbDirFlag,
		dbNameFlag,
	},
	SkipFlagParsing:        false,
	HideHelp:               false,
	HideHelpCommand:        false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}
