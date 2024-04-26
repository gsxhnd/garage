package task

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gsxhnd/garage/garage_ffmpeg"
	"github.com/reactivex/rxgo/v2"
)

type ffmpegTask struct {
	id      string
	cmd     string
	batcher garage_ffmpeg.VideoBatcher
	ob      *taskOb
}

func NewFFmpegTask(opt *garage_ffmpeg.VideoBatchOption, cmd string) (Task, error) {
	batcher, err := garage_ffmpeg.NewVideoBatch(opt)
	if err != nil {
		return nil, err
	}

	return &ffmpegTask{
		id:      uuid.New().String(),
		cmd:     cmd,
		batcher: batcher,
		ob:      newTaskOb(),
	}, nil
}

func (t *ffmpegTask) GetId() string {
	return t.id
}

func (t *ffmpegTask) GetOB() rxgo.Observable {
	return t.ob.ob
}

func (t *ffmpegTask) Run() {
	switch t.cmd {
	case "convert":
		cmds, err := t.batcher.GetConvertBatch()
		for _, c := range cmds {
			fmt.Println("cmd: ", c)
		}

		if err != nil || len(cmds) == 0 {
			return
		}
		t.ob.ch <- rxgo.Of(cmds)
		if err := t.batcher.ExecuteBatch(t.ob, t.ob, cmds); err != nil {
			t.ob.ch <- rxgo.Of("error executing batch")
			close(t.ob.ch)
			return
		}
	default:
	}
}
