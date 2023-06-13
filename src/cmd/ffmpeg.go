package cmd

import (
	"os"
	"os/exec"

	"github.com/gsxhnd/garage/src/batch"
	"github.com/gsxhnd/garage/src/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	ffmpegBatchInputPathFlag = &cli.StringFlag{
		Name:  "input-path",
		Value: "./",
		Usage: "源视频路径",
	}
	ffmpegBatchInputTypeFlag = &cli.StringFlag{
		Name:  "input-type",
		Value: ".mkv",
		Usage: "源视频后缀",
	}
	ffmpegBatchInputSubSuffixFlag = &cli.StringFlag{
		Name:  "input-sub-suffix",
		Value: ".ass",
		Usage: "添加的字幕后缀",
	}
	ffmpegBatchInputSubNoFlag = &cli.IntFlag{
		Name:  "input-sub-no",
		Value: 0,
		Usage: "添加的字幕所处流的位置",
	}
	ffmpegBatchInputSubLangFlag = &cli.StringFlag{
		Name:  "input-sub-lang",
		Value: "chi",
		Usage: "添加的字幕语言缩写其他语言请参考ffmpeg",
	}
	ffmpegBatchInputSubTitleFlag = &cli.StringFlag{
		Name:  "input-sub-title",
		Value: "Chinese",
		Usage: "添加的字幕标题",
	}
	ffmpegBatchInputFontsPathFlag = &cli.StringFlag{
		Name:  "input-fonts-path",
		Usage: "添加的字体文件夹",
	}
	ffmpegBatchOutputDestPathFlag = &cli.StringFlag{
		Name:  "output-path",
		Value: "./result/",
		Usage: "转换后文件存储位置",
	}
	ffmpegBatchOutputTypeFlag = &cli.StringFlag{
		Name:  "output-type",
		Value: ".mkv",
		Usage: "转换后的视频后缀",
	}
	ffmpegBatchAdvanceFlag = &cli.StringFlag{
		Name:  "advance",
		Value: "",
		Usage: "高级自定义参数",
	}
	ffmpegBatchExecFlag = &cli.BoolFlag{
		Name:  "exec",
		Value: false,
		Usage: "是否执行批处理命令False时仅打印命令",
	}
)

var ffmpegBatchCmd = &cli.Command{
	Name: "ffmpeg-batch",
	Subcommands: []*cli.Command{
		ffmpegBatchConvertCmd,
		ffmpegBatchAddSubCmd,
	},
}

var ffmpegBatchConvertCmd = &cli.Command{
	Name:      "convert",
	Aliases:   nil,
	Usage:     "视频转换批处理",
	UsageText: "",
	Flags: []cli.Flag{
		ffmpegBatchInputPathFlag,
		ffmpegBatchInputTypeFlag,
		ffmpegBatchOutputDestPathFlag,
		ffmpegBatchOutputTypeFlag,
		ffmpegBatchAdvanceFlag,
		ffmpegBatchExecFlag,
	},
	Action: func(c *cli.Context) error {
		var (
			logger     = utils.GetLogger()
			inputPath  = c.String("input-path")
			inputType  = c.String("input-type")
			outputPath = c.String("output-path")
			outputType = c.String("output-type")
			advance    = c.String("advance")
		)
		vb, err := batch.NewVideoBatch(logger, inputPath, inputType, outputPath)
		if err != nil {
			logger.Panic("Create dest path error", zap.Error(err))
			return err
		}

		batch, err := vb.GetConvertBatch(advance, outputType)
		if err != nil {
			logger.Panic("Get convert cmd batch error", zap.Error(err))
		}

		logger.Info("Get all videos, starting convert")
		for _, cmd := range batch {
			if !c.Bool("exec") {
				logger.Sugar().Infof("cmd: %v", cmd)
			} else {
				logger.Sugar().Infof("Start convert video cmd: %v", cmd)
				cmd := exec.Command("powershell", cmd)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					logger.Sugar().Errorf("cmd errror: %v", err)
				}
				logger.Sugar().Infof("Finished convert video")
			}
		}
		return nil
	},
}

var ffmpegBatchAddSubCmd = &cli.Command{
	Name:      "add-sub",
	Aliases:   nil,
	Usage:     "视频添加字幕批处理",
	UsageText: "",
	Flags: []cli.Flag{
		ffmpegBatchInputPathFlag,
		ffmpegBatchInputTypeFlag,
		ffmpegBatchInputSubSuffixFlag,
		ffmpegBatchInputSubNoFlag,
		ffmpegBatchInputSubTitleFlag,
		ffmpegBatchInputSubLangFlag,
		ffmpegBatchInputFontsPathFlag,
		ffmpegBatchOutputDestPathFlag,
		ffmpegBatchExecFlag,
	},
	Action: func(c *cli.Context) error {
		var (
			logger         = utils.GetLogger()
			inputPath      = c.String("input-path")
			inputType      = c.String("input-type")
			inputSubNo     = c.Int("input-sub-no")
			inputSubSuffix = c.String("input-sub-suffix")
			inputSubLang   = c.String("input-sub-lang")
			inputSubTitle  = c.String("input-sub-title")
			inputFontsPath = c.String("input-fonts-path")
			outputPath     = c.String("output-path")
		)
		vb, err := batch.NewVideoBatch(logger, inputPath, inputType, outputPath)
		if err != nil {
			logger.Panic("Create dest path error", zap.Error(err))
			return err
		}

		batch, err := vb.GetAddSubtitleBatch(
			inputSubNo,
			inputSubSuffix,
			inputSubLang,
			inputSubTitle,
			inputFontsPath,
		)
		if err != nil {
			logger.Panic("Get add subtitle batch error", zap.Error(err))
		}
		logger.Info("Get all videos, starting add subtitle")
		for _, cmd := range batch {
			if !c.Bool("exec") {
				logger.Sugar().Infof("cmd: %v", cmd)
			} else {
				logger.Sugar().Infof("Start add subtitle into video, cmd: %v", cmd)
				cmd := exec.Command("powershell", cmd)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					logger.Sugar().Errorf("cmd errror: %v", err)
				}
				logger.Sugar().Infof("Finished  add subtitle into video")
			}
		}
		return nil
	},
}
