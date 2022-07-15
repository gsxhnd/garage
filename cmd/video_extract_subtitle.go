package cmd

import (
	"github.com/urfave/cli/v2"
)

var videoExtractSubtitleCmd = &cli.Command{
	Name:      "video_convert",
	Aliases:   nil,
	Usage:     "视频提取字幕",
	UsageText: "",
	Flags: []cli.Flag{
		exec_flag,
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
