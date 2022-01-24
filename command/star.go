package command

import "github.com/urfave/cli/v2"

var starCmd = &cli.Command{
	Name:  "star",
	Usage: "crawl jav data by star code",
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) error {
		return nil 
	},
}
