package cmd

import "github.com/urfave/cli/v2"

var javPrefixCmd = &cli.Command{
	Name:  "jav_prefix",
	Usage: "根据番号前缀爬取数据",
	Flags: []cli.Flag{
		proxyFlag,
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
