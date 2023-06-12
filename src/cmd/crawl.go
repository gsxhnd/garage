package cmd

import (
	"github.com/urfave/cli/v2"
)

var crawlCmd = &cli.Command{
	Name: "crawl",
	Subcommands: []*cli.Command{
		javCodeCmd,
		javPrefixCmd,
		javStarCodeCmd,
	},
}
