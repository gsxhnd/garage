package cmd

import (
	"strconv"

	"github.com/gsxhnd/garage/batch"
	"github.com/gsxhnd/garage/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var videoSubtitleCmd = &cli.Command{
	Name:      "video_subtitle",
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
		fonts_flag,
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
			DestPath:               c.String("dest_path"),
			DestVideoType:          c.String("dest_video_type"),
			Advance:                c.String("advance"),
			Logger:                 logger,
		}
		fonts := c.StringSlice("fonts")

		var (
			cmd  = make(chan string)
			done = make(chan bool)
		)
		go vb.CreateBatchFile(cmd, done)

		vl, err := vb.GetVideos()
		if err != nil {
			vb.Logger.Error("Get videos error", zap.Error(err))
			return nil
		}
		vb.Logger.Info("Get all videos, starting add subtitle")
		if err := vb.CreateDestDir(); err != nil {
			return err
		}
		vb.Logger.Info("Source videos directory: " + vb.SourceRootPath)
		vb.Logger.Info("Get matching video count: " + strconv.Itoa(len(vl)))
		vb.Logger.Info("Target video's subtitle stream number: " + strconv.Itoa(vb.SourceSubtitleNumber))
		vb.Logger.Info("Target video's subtitle language: " + vb.SourceSubtitleLanguage)
		vb.Logger.Info("Target video's subtitle title: " + vb.SourceSubtitleTitle)

		for _, v := range fonts {
			vb.Logger.Info("Target video's subtitlle fonts: " + v)
		}
		vb.Logger.Info("Dest video directory: " + vb.DestPath)

		batch := vb.GetSubtitleBatch(vl)

		for _, v := range batch {
			cmd <- v
		}
		close(cmd)
		d := <-done
		if d {
			vb.Logger.Info("Write batch file complete...")
		} else {
			vb.Logger.Info("Write batch file failed...")
		}
		return nil
	},
}
