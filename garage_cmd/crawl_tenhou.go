package main

import "github.com/urfave/cli/v2"

var crawlTenhouCmd = &cli.Command{
	Name:  "crawl_tenhou",
	Usage: "",
	Action: func(c *cli.Context) error {
		return nil
	},
}
