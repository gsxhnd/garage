package cmd

import (
	"github.com/urfave/cli/v2"
)

var (
	proxyFlag = &cli.StringFlag{Name: "proxy", Usage: "代理配置,如: http://127.0.0.1:1080"}
	siteFlag  = &cli.StringFlag{
		Name:        "site",
		Usage:       "选择爬取数据的网站,支持网站(javbus)",
		Destination: nil,
		HasBeenSet:  false,
		Value:       "javbus",
	}
)

var (
	source_root_path_flag = &cli.StringFlag{
		Name:  "source_root_path",
		Value: "./",
	}
	source_video_type_flag = &cli.StringFlag{
		Name:  "source_video_type",
		Value: ".mkv",
	}
	source_subtitle_type_flag = &cli.StringFlag{
		Name:  "source_subtitle_type",
		Value: ".ass",
	}
	source_subtitle_number_flag = &cli.IntFlag{
		Name:  "source_subtitle_number_type",
		Value: 0,
	}
	source_subtitle_language_flag = &cli.StringFlag{
		Name:  "source_subtitle_language",
		Value: "chi",
	}
	source_subtitle_title_flag = &cli.StringFlag{
		Name:  "source_subtitle_title",
		Value: "Chinese",
	}
	dest_path_flag = &cli.StringFlag{
		Name:  "dest_path",
		Value: "./result/",
	}
	dest_video_type_flag = &cli.StringFlag{
		Name:  "dest_video_type",
		Value: ".mkv",
	}
	exec_flag = &cli.BoolFlag{
		Name:  "exec",
		Value: true,
	}
	advance_flag = &cli.StringFlag{
		Name:  "advance",
		Value: "",
	}
)
