package command

import "github.com/urfave/cli/v2"

var (
	portFlag = &cli.StringFlag{
		Name:        "port",
		Aliases:     []string{"p"},
		Usage:       "",
		Destination: nil,
		HasBeenSet:  false,
		Value:       "9000",
	}
	coverImgFlag = &cli.StringFlag{
		Name:    "cover-image",
		Aliases: []string{"ci"},
	}
	starImgFlag = &cli.StringFlag{
		Name:    "star-image",
		Aliases: []string{"si"},
	}
)
