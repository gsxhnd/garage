package command

import (
	"garage/api"
	"github.com/urfave/cli/v2"
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
	Before:       nil,
	After:        nil,
	Action: func(context *cli.Context) error {
		go api.OpenBrowser("http://localhost:8001")
		api.Run()
		return nil
	},
	OnUsageError:           nil,
	Flags:                  nil,
	SkipFlagParsing:        false,
	HideHelp:               false,
	HideHelpCommand:        false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}
