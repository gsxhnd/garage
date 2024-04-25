package task

import (
	"fmt"

	"github.com/gsxhnd/garage/garage_ffmpeg"
)

type ffmpegTask struct {
	id      string
	batcher garage_ffmpeg.VideoBatcher
}

func NewFFmpegTask(opt *garage_ffmpeg.VideoBatchOption) (Task, error) {
	batcher, err := garage_ffmpeg.NewVideoBatch(opt)
	if err != nil {
		return nil, err
	}

	return &ffmpegTask{
		id:      "",
		batcher: batcher,
	}, nil
}

func (t *ffmpegTask) GetId() string {
	return t.id
}

func (t *ffmpegTask) Run(cmd string) {
	switch cmd {
	case "convert":
		fmt.Println("convert")
	default:
	}
}
