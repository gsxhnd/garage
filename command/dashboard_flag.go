package command

import "github.com/urfave/cli/v2"

var (
	confFlag = &cli.StringFlag{
		Name:        "conf",
		Aliases:     []string{"c"},
		Usage:       "-c [Conf File]",
		Destination: nil,
		HasBeenSet:  false,
		Value:       "./conf.yaml",
	}
	openFlag = &cli.BoolFlag{
		Name:        "open",
		Usage:       "--open",
		Destination: nil,
		Value:       false,
	}
)
