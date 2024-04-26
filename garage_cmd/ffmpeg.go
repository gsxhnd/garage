package main

import (
	"os"

	"github.com/gsxhnd/garage/garage_ffmpeg"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	inputPath = &cli.StringFlag{
		Name:  "input_path",
		Value: "./",
		Usage: "源视频路径",
	}

	inputFormat = &cli.StringFlag{
		Name:  "input_format",
		Value: "mkv",
		Usage: "源视频后缀",
	}

	outputPath = &cli.StringFlag{
		Name:  "output_path",
		Value: "./result/",
		Usage: "转换后文件存储位置",
	}

	outputFormat = &cli.StringFlag{
		Name:  "output_format",
		Value: "mkv",
		Usage: "转换后的视频后缀",
	}

	advance = &cli.StringFlag{
		Name:  "advance",
		Value: "",
		Usage: "高级自定义参数",
	}

	exec = &cli.BoolFlag{
		Name:  "exec",
		Value: false,
		Usage: "是否执行批处理命令False时仅打印命令",
	}

	inputFontsPath = &cli.StringFlag{
		Name:     "input_fonts_path",
		Usage:    "添加的字体文件夹",
		Required: false,
	}
)

var ffmpegBatchCmd = &cli.Command{
	Name:        "ffmpeg_batch",
	Description: "ffmpeg视频批处理工具，支持视频格式转换、字幕添加和字体添加",
	Flags:       []cli.Flag{},
	Subcommands: []*cli.Command{
		ffmpegBatchConvertCmd,
		ffmpegBatchAddSubCmd,
		ffmpegBatchAddFontCmd,
	},
}

var ffmpegBatchConvertCmd = &cli.Command{
	Name:    "convert",
	Aliases: nil,
	Flags: []cli.Flag{
		inputPath,
		inputFormat,
		outputPath,
		outputFormat,
		advance,
		exec,
	},
	Usage:     "视频转换批处理",
	UsageText: "",
	Action: func(ctx *cli.Context) error {
		var opt = garage_ffmpeg.VideoBatchOption{
			InputPath:    ctx.String("input_path"),
			InputFormat:  ctx.String("input_format"),
			OutputPath:   ctx.String("output_path"),
			OutputFormat: ctx.String("output_format"),
			Advance:      ctx.String("advance"),
			Exec:         ctx.Bool("exec"),
		}
		logger.Infof("Source videos directory: " + opt.InputPath)
		logger.Infof("Source videos format: " + opt.InputFormat)
		logger.Infof("Target video's font paths: " + opt.FontsPath)
		logger.Infof("Dest video directory: " + opt.OutputPath)
		logger.Infof("Dest video format: " + opt.OutputFormat)

		vb, err := garage_ffmpeg.NewVideoBatch(&opt)
		if err != nil {
			logger.Panicf("Create dest path error", zap.Error(err))
			return err
		}
		logger.Debugf("video batcher init")

		cmds, err := vb.GetConvertBatch()
		if err != nil {
			logger.Panicf("Get Convert Batch error", zap.Error(err))
			return err
		}

		if !opt.Exec {
			for _, cmd := range cmds {
				for _, c := range cmd {
					logger.Infof("Cmd batch not execute,cmd: " + c)
				}
			}
			return nil
		}

		return vb.ExecuteBatch(os.Stdout, os.Stderr, cmds)
	},
}

var ffmpegBatchAddSubCmd = &cli.Command{
	Name:      "add_sub",
	Aliases:   nil,
	Usage:     "视频添加字幕批处理",
	UsageText: "",
	Flags: []cli.Flag{
		inputPath,
		inputFormat,
		outputPath,
		outputFormat,
		advance,
		inputFontsPath,
		exec,
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
		var opt = garage_ffmpeg.VideoBatchOption{
			InputPath:    ctx.String("input_path"),
			InputFormat:  ctx.String("input_format"),
			OutputPath:   ctx.String("output_path"),
			OutputFormat: ctx.String("output_format"),
			Advance:      ctx.String("advance"),
			Exec:         ctx.Bool("exec"),
			FontsPath:    ctx.String("input_fonts_path"),
		}

		vb, err := garage_ffmpeg.NewVideoBatch(&opt)
		if err != nil {
			logger.Panicf("Create dest path error", zap.Error(err))
			return err
		}

		// vb.logger.Debug("Get matching video count: " + strconv.Itoa(len(vb.videosList)))
		// vb.logger.Debug("Target video's subtitle stream number: " + strconv.Itoa(vb.option.InputSubNo))
		// vb.logger.Debug("Target video's subtitle language: " + vb.option.InputSubLang)
		// vb.logger.Debug("Target video's subtitle title: " + vb.option.InputSubTitle)
		// vb.logger.Info("Target video's font paths not set, skip.")
		logger.Infof("Target video's font paths: " + opt.FontsPath)

		cmds, err := vb.GetAddSubtittleBatch()
		if err != nil {
			return err
		}

		if !opt.Exec {
			for _, cmd := range cmds {
				logger.Infof(cmd)
			}
			return nil
		}
		return nil

		// return vb.ExecuteBatch(os.Stdout, os.Stderr, cmds)
	},
}

var ffmpegBatchAddFontCmd = &cli.Command{
	Name:    "add_fonts",
	Aliases: nil,
	Usage:   "视频添加字体批处理",
	Flags: []cli.Flag{
		inputPath,
		inputFormat,
		outputPath,
		outputFormat,
		advance,
		exec,
		inputFontsPath,
	},
	UsageText: "",
	Action: func(ctx *cli.Context) error {
		var opt = garage_ffmpeg.VideoBatchOption{
			InputPath:    ctx.String("input_path"),
			InputFormat:  ctx.String("input_format"),
			OutputPath:   ctx.String("output_path"),
			OutputFormat: ctx.String("output_format"),
			Advance:      ctx.String("advance"),
			Exec:         ctx.Bool("exec"),
			FontsPath:    ctx.String("input_fonts_path"),
		}

		vb, err := garage_ffmpeg.NewVideoBatch(&opt)
		if err != nil {
			logger.Panicf("Create dest path error", zap.Error(err))
			return err
		}
		cmds, err := vb.GetAddFontsBatch()
		if err != nil {
			return err
		}

		if !opt.Exec {
			for _, cmd := range cmds {
				logger.Infof(cmd)
			}
			return nil
		}
		return nil

		// return vb.ExecuteBatch(os.Stdout, os.Stderr, cmds)
	},
}
