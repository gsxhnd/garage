package cmd

import (
	"github.com/gsxhnd/garage/src/cmd/crawl_cmd"
	"github.com/urfave/cli/v2"
)

var crawlCmd = &cli.Command{
	Name: "crawl",
	Subcommands: []*cli.Command{
		crawl_cmd.JavCodeCmd,
		crawl_cmd.JavPrefixCmd,
		crawl_cmd.JavStarCodeCmd,
	},
}
