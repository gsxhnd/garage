package command

import "github.com/urfave/cli/v2"

var starCmd = &cli.Command{
	Name:  "star",
	Usage: "根据演员ID爬取数据",
	Flags: []cli.Flag{
		proxyFlag,
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
