package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gsxhnd/garage/batch"
	"github.com/gsxhnd/garage/utils"
	"github.com/urfave/cli/v2"
)

var videoConvertCmd = &cli.Command{
	Name:      "video_convert",
	Aliases:   nil,
	Usage:     "",
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
		vl, _ := vb.GetVideos()
		batch := vb.GetConvertBatch(vl)
		if c.Bool("exec") {
			execShell(batch)
		} else {
			for _, v := range batch {
				fmt.Println(v)
			}
		}
		return nil
	},
}

func execShell(cmd []string) {
	for _, v := range cmd {
		fmt.Println(v)
		c := exec.Command("powershell", v)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	}
}
