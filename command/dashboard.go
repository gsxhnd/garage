package command

import "github.com/urfave/cli/v2"

// start dashboard api
var dashboardCmd = &cli.Command{
	Name:                   "dashboard",
	Aliases:                nil,
	Usage:                  "",
	UsageText:              "",
	Description:            "",
	ArgsUsage:              "",
	Category:               "",
	BashComplete:           nil,
	Before:                 nil,
	After:                  nil,
	Action:                 nil,
	OnUsageError:           nil,
	Subcommands:            nil,
	Flags:                  nil,
	SkipFlagParsing:        false,
	HideHelp:               false,
	HideHelpCommand:        false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}
