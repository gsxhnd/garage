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
		dest_video_type_flag,
		advance_flag,
		exec_flag,
	},
	Action: func(c *cli.Context) error {
		logger := utils.GetLogger()
		var vb = &batch.VideoBatch{
			SourceRootPath:         c.String("source_root_path"),
			SourceVideoType:        c.String("source_video_type"),
			SourceSubtitleType:     c.String("source_subtitle_type"),
			SourceSubtitleNumber:   c.Int("source_subtitle_number"),
			SourceSubtitleLanguage: c.String("source_subtitle_language"),
			SourceSubtitleTitle:    c.String("source_subtitle_title"),
			FontsPath:              c.String("fonts_path"),
			DestPath:               c.String("dest_path"),
			DestVideoType:          c.String("dest_video_type"),
			Advance:                c.String("advance"),
			Logger:                 logger,
		}

		vl, err := vb.GetVideos()
		if err != nil {
			vb.Logger.Error("Get videos error", zap.Error(err))
			return err
		}
		vb.Logger.Info("Get all videos, starting add subtitle")
		if err := vb.CreateDestDir(); err != nil {
			return err
		}

		batch := vb.GetImportSubtitleBatch(vl)
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
