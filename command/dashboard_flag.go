package command

import "github.com/urfave/cli/v2"

var (
	openFlag = &cli.BoolFlag{
		Name:        "open",
		Usage:       "--open",
		Destination: nil,
		Value:       false,
	}
)
