package command

import "github.com/urfave/cli/v2"

var codeCommand = &cli.Command{
	Name: "code",
	Action: func(c *cli.Context) error {
		return nil
	},
}
