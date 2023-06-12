package crawl_cmd

import "github.com/urfave/cli/v2"

var javStarCmd = &cli.Command{
	Name:  "jav_star_code",
	Usage: "根据演员ID爬取数据",
	Flags: []cli.Flag{
		siteFlag,
		proxyFlag,
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
