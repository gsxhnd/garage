package main

import (
	"context"

	"github.com/gsxhnd/garage/garage_ffmpeg"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

type garage_ffmpeg_option string

var ffmpegBatchCmd = &cli.Command{
	Name:        "ffmpeg_batch",
	Description: "ffmpeg视频批处理工具，支持视频格式转换、字幕添加和字体添加",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "input_path",
			Value: "./",
			Usage: "源视频路径",
		}, &cli.StringFlag{
			Name:  "input_format",
			Value: ".mkv",
			Usage: "源视频后缀",
		},
		&cli.StringFlag{
			Name:  "output_path",
			Value: "./result/",
			Usage: "转换后文件存储位置",
		},
		&cli.StringFlag{
			Name:  "output_format",
			Value: ".mkv",
			Usage: "转换后的视频后缀",
		},
		&cli.StringFlag{
			Name:  "advance",
			Value: "",
			Usage: "高级自定义参数",
		},
		&cli.BoolFlag{
			Name:  "exec",
			Value: false,
			Usage: "是否执行批处理命令False时仅打印命令",
		},
	},
	Subcommands: []*cli.Command{
		ffmpegBatchConvertCmd,
		ffmpegBatchAddSubCmd,
		ffmpegBatchAddFontCmd,
	},
	Before: func(ctx *cli.Context) error {
		var opt = garage_ffmpeg.VideoBatchOption{
			InputPath:    ctx.String("input_path"),
			InputFormat:  ctx.String("input_format"),
			OutputPath:   ctx.String("output_path"),
			OutputFormat: ctx.String("output_format"),
			Advance:      ctx.String("advance"),
			Exec:         ctx.Bool("exec"),
		}
		ctx.Context = context.WithValue(ctx.Context, garage_ffmpeg_option("opt"), opt)
		return nil
	},
}

var ffmpegBatchConvertCmd = &cli.Command{
	Name:      "convert",
	Aliases:   nil,
	Usage:     "视频转换批处理",
	UsageText: "",
	Action: func(c *cli.Context) error {
		var (
			opt = c.Context.Value(garage_ffmpeg_option("opt")).(garage_ffmpeg.VideoBatchOption)
		)
		vb, err := garage_ffmpeg.NewVideoBatch(&opt)
		if err != nil {
			logger.Panicf("Create dest path error", zap.Error(err))
			return err
		}
		return vb.StartConvertBatch()
	},
}

var ffmpegBatchAddSubCmd = &cli.Command{
	Name:      "add_sub",
	Aliases:   nil,
	Usage:     "视频添加字幕批处理",
	UsageText: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "input_fonts_path",
			Usage: "添加的字体文件夹",
		},
		&cli.StringFlag{
			Name:  "input_sub_suffix",
			Value: ".ass",
			Usage: "添加的字幕后缀",
		},
		&cli.IntFlag{
			Name:  "input_sub_no",
			Value: 0,
			Usage: "添加的字幕所处流的位置",
		},
		&cli.StringFlag{
			Name:  "input_sub_lang",
			Value: "chi",
			Usage: "添加的字幕语言缩写其他语言请参考ffmpeg",
		},
		&cli.StringFlag{
			Name:  "input_sub_title",
			Value: "Chinese",
			Usage: "添加的字幕标题",
		},
	},
	Action: func(ctx *cli.Context) error {
		var opt = ctx.Context.Value(garage_ffmpeg_option("opt")).(garage_ffmpeg.VideoBatchOption)
		opt.FontsPath = ctx.String("input_fonts_path")
		vb, err := garage_ffmpeg.NewVideoBatch(&opt)
		if err != nil {
			logger.Panicf("Create dest path error", zap.Error(err))
			return err
		}

		return vb.StartAddSubtittleBatch()
	},
}

var ffmpegBatchAddFontCmd = &cli.Command{
	Name:      "add_fonts",
	Aliases:   nil,
	Usage:     "视频添加字体批处理",
	UsageText: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "input_fonts_path",
			Usage:    "添加的字体文件夹",
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		var opt = ctx.Context.Value(garage_ffmpeg_option("opt")).(garage_ffmpeg.VideoBatchOption)
		opt.FontsPath = ctx.String("input_fonts_path")
		vb, err := garage_ffmpeg.NewVideoBatch(&opt)

		if err != nil {
			logger.Panicf("Create dest path error", zap.Error(err))
			return err
		}

		return vb.StartAddFontsBatch()
		// if err != nil {
		// 	logger.Panic("Get add subtitle batch error", zap.Error(err))
		// }
		// logger.Info("Get all videos, starting add fonts")
		// for _, cmd := range batch {
		// 	if !c.Bool("exec") {
		// 		logger.Sugar().Infof("cmd: %v", cmd)
		// 	} else {
		// 		startTime := time.Now()
		// 		logger.Sugar().Infof("Start add subtitle into video, cmd: %v", cmd)
		// 		cmd := exec.Command("powershell", cmd)
		// 		cmd.Stdout = os.Stdout
		// 		cmd.Stderr = os.Stderr
		// 		err := cmd.Run()
		// 		if err != nil {
		// 			logger.Sugar().Errorf("cmd errror: %v", err)
		// 		}
		// 		logger.Sugar().Infof("Finished  add subtitle into video")
		// 		logger.Sugar().Infof("Finished convert video, spent time: %v sec", time.Since(startTime).Seconds())
		// 	}
		// }
	},
}
