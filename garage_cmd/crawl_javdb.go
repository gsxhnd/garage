package main

import "github.com/urfave/cli/v2"

var crawlJavDBCmd = &cli.Command{
	Name:  "crawl_javdb",
	Usage: "",
	Action: func(c *cli.Context) error {
		return nil
	},
}
