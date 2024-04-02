package task

import "github.com/gsxhnd/garage/garage_ffmpeg"

type ffmpegTask struct {
	batcher garage_ffmpeg.VideoBatcher
}

func NewFFmpegTask(batcher garage_ffmpeg.VideoBatcher) (Task, error) {
	return &ffmpegTask{
		batcher: batcher,
	}, nil
}
