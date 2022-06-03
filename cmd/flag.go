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
		Usage: "源视频路径",
	}
	source_video_type_flag = &cli.StringFlag{
		Name:  "source_video_type",
		Value: ".mkv",
		Usage: "源视频后缀",
	}
	source_subtitle_type_flag = &cli.StringFlag{
		Name:  "source_subtitle_type",
		Value: ".ass",
		Usage: "添加的字幕后缀",
	}
	source_subtitle_number_flag = &cli.IntFlag{
		Name:  "source_subtitle_number",
		Value: 0,
		Usage: "添加的字幕所处流的位置",
	}
	source_subtitle_language_flag = &cli.StringFlag{
		Name:  "source_subtitle_language",
		Value: "chi",
		Usage: "添加的字幕语言缩写，其他语言请参考ffmpeg",
	}
	source_subtitle_title_flag = &cli.StringFlag{
		Name:  "source_subtitle_title",
		Value: "Chinese",
		Usage: "添加的字幕标题",
	}
	dest_path_flag = &cli.StringFlag{
		Name:  "dest_path",
		Value: "./result/",
		Usage: "转换后文件存储位置",
	}
	dest_video_type_flag = &cli.StringFlag{
		Name:  "dest_video_type",
		Value: ".mkv",
		Usage: "转换后的视频后缀",
	}
	exec_flag = &cli.BoolFlag{
		Name:  "exec",
		Value: true,
		Usage: "是否执行批处理命令，False时仅打印命令",
	}
	advance_flag = &cli.StringFlag{
		Name:  "advance",
		Value: "",
		Usage: "高级自定义参数",
	}
)
