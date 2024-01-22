package cmd

import (
	"github.com/gsxhnd/garage/ffmpeg"
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
	Name:        "ffmpeg-batch",
	Description: "ffmpeg视频批处理工具，支持视频格式转换、字幕添加和字体添加",
	Subcommands: []*cli.Command{
		ffmpegBatchConvertCmd,
		ffmpegBatchAddSubCmd,
		ffmpegBatchAddFontCmd,
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
			logger = utils.GetLogger()
			// inputPath  = c.String("input-path")
			// inputType  = c.String("input-type")
			// outputPath = c.String("output-path")
			// outputType = c.String("output-type")
			// advance    = c.String("advance")
		)

		vb, err := ffmpeg.NewVideoBatch(logger, nil)
		if err != nil {
			logger.Panic("Create dest path error", zap.Error(err))
			return err
		}
		return vb.StartConvertBatch()
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
			logger = utils.GetLogger()
			// inputPath      = c.String("input-path")
			// inputType      = c.String("input-type")
			// inputSubNo     = c.Int("input-sub-no")
			// inputSubSuffix = c.String("input-sub-suffix")
			// inputSubLang   = c.String("input-sub-lang")
			// inputSubTitle  = c.String("input-sub-title")
			// inputFontsPath = c.String("input-fonts-path")
			// outputPath     = c.String("output-path")
		)
		vb, err := ffmpeg.NewVideoBatch(logger, nil)
		if err != nil {
			logger.Panic("Create dest path error", zap.Error(err))
			return err
		}

		return vb.StartAddSubtittleBatch()
		// if err != nil {
		// 	logger.Panic("Get add subtitle batch error", zap.Error(err))
		// }
		// logger.Info("Get all videos, starting add subtitle")
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

var ffmpegBatchAddFontCmd = &cli.Command{
	Name:      "add-fonts",
	Aliases:   nil,
	Usage:     "视频添加字体批处理",
	UsageText: "",
	Flags: []cli.Flag{
		ffmpegBatchInputPathFlag,
		ffmpegBatchInputTypeFlag,
		ffmpegBatchInputFontsPathFlag,
		ffmpegBatchOutputDestPathFlag,
		ffmpegBatchExecFlag,
	},
	Action: func(c *cli.Context) error {
		var (
			logger = utils.GetLogger()
			// inputPath      = c.String("input-path")
			// inputType      = c.String("input-type")
			// inputFontsPath = c.String("input-fonts-path")
			// outputPath     = c.String("output-path")
		)
		vb, err := ffmpeg.NewVideoBatch(logger, nil)
		if err != nil {
			logger.Panic("Create dest path error", zap.Error(err))
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
