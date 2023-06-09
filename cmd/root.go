package cmd

import (
	"github.com/gsxhnd/garage/cmd/ffmpeg_cmd"
	"github.com/gsxhnd/garage/cmd/jav_cmd"
	"github.com/urfave/cli/v2"
)

var (
	RootCmd = cli.NewApp()
)

func init() {
	RootCmd.HideVersion = true
	RootCmd.Usage = "命令行工具"
	RootCmd.Flags = []cli.Flag{}
	RootCmd.Commands = []*cli.Command{
		jav_cmd.CodeCmd,
		ffmpeg_cmd.VideoConvertCmd,
		ffmpeg_cmd.VideoSubtitleCmd,
		versionCmd,
	}
}
