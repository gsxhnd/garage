package ffmpeg_cmd

import (
	"os"
	"os/exec"

	"github.com/gsxhnd/garage/src/batch"
	"github.com/gsxhnd/garage/src/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var VideoConvertCmd = &cli.Command{
	Name:      "video_convert",
	Aliases:   nil,
	Usage:     "视频转换批处理",
	UsageText: "",
	Flags: []cli.Flag{
		source_root_path_flag,
		source_video_type_flag,
		dest_path_flag,
		dest_video_type_flag,
		advance_flag,
		exec_flag,
	},
	Action: func(c *cli.Context) error {
		logger := utils.GetLogger()
		vb := batch.NewVideoBatch(logger, c.String("source_root_path"), c.String("source_video_type"))
		if err := vb.CreateDestDir(c.String("dest_path")); err != nil {
			logger.Panic("Create dest path error", zap.Error(err))
			return err
		}

		batch, err := vb.GetConvertBatch(c.String("advance"), c.String("dest_video_type"))
		if err != nil {
			logger.Panic("Get convert cmd batch error", zap.Error(err))
		}

		logger.Info("Get all videos, starting convert")
		for _, cmd := range batch {
			if !c.Bool("exec") {
				logger.Sugar().Infof("cmd: %v", cmd)
			} else {
				cmd := exec.Command("powershell", cmd)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					logger.Sugar().Errorf("cmd errror: %v", err)
				}
			}
		}
		return nil
	},
}
