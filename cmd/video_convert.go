package cmd

import (
	"github.com/gsxhnd/garage/batch"
	"github.com/gsxhnd/garage/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var videoConvertCmd = &cli.Command{
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

		vb.Logger.Info("Get all videos, starting convert")
		if err := vb.CreateDestDir(); err != nil {
			return err
		}
		batch := vb.GetConvertBatch(vl)
		if err := vb.CreateDestDir(); err != nil {
			return err
		}
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
