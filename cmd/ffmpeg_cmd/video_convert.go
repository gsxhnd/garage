package ffmpeg_cmd

import (
	"os"
	"os/exec"

	"github.com/gsxhnd/garage/batch"
	"github.com/gsxhnd/garage/utils"
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
		var vb = batch.VideoBatch{
			SourceRootPath:  c.String("source_root_path"),
			SourceVideoType: c.String("source_video_type"),
			DestPath:        c.String("dest_path"),
			DestVideoType:   c.String("dest_video_type"),
			Advance:         c.String("advance"),
			Logger:          logger,
		}

		vl, err := vb.GetVideos()
		if err != nil {
			vb.Logger.Error("Get videos error", zap.Error(err))
			return nil
		}

		vb.Logger.Info("Get all videos, starting convert")
		if err := vb.CreateDestDir(); err != nil {
			return err
		}
		if err := vb.CreateDestDir(); err != nil {
			return err
		}
		batch := vb.GetConvertBatch(vl)
		for _, cmd := range batch {
			if !c.Bool("exec") {
				vb.Logger.Sugar().Infof("cmd: %v", cmd)
			} else {
				cmd := exec.Command("powershell", cmd)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					vb.Logger.Sugar().Errorf("cmd errror: %v", err)
				}
			}
		}
		return nil
	},
}
