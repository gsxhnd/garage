package ffmpeg_cmd

import (
	"os"
	"os/exec"

	"github.com/gsxhnd/garage/src/batch"
	"github.com/gsxhnd/garage/src/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var VideoSubtitleCmd = &cli.Command{
	Name:      "video_import_subtitle",
	Aliases:   nil,
	Usage:     "视频添加字幕批处理",
	UsageText: "",
	Flags: []cli.Flag{
		source_root_path_flag,
		source_video_type_flag,
		source_subtitle_type_flag,
		source_subtitle_number_flag,
		source_subtitle_language_flag,
		source_subtitle_title_flag,
		fonts_path_flag,
		dest_path_flag,
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

		batch, err := vb.GetAddSubtitleBatch(
			c.Int("source_subtitle_number"),
			c.String("source_subtitle_type"),
			c.String("source_subtitle_language"),
			c.String("source_subtitle_title"),
			c.String("fonts_path"),
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
