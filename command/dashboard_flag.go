package command

import "github.com/urfave/cli/v2"

var (
	portFlag = &cli.StringFlag{
		Name:        "port",
		Aliases:     []string{"p"},
		Usage:       "",
		Destination: nil,
		HasBeenSet:  false,
	}
)
