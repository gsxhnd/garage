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
	exec    bool
	cmds    [][]string
}

func NewFFmpegTask(opt *garage_ffmpeg.VideoBatchOption, cmd string) (Task, error) {
	batcher, err := garage_ffmpeg.NewVideoBatch(opt)
	if err != nil {
		return nil, err
	}

	var task = &ffmpegTask{
		id:      uuid.New().String(),
		cmd:     cmd,
		batcher: batcher,
		exec:    opt.Exec,
		ob:      newTaskOb(),
		cmds:    make([][]string, 0),
	}

	switch cmd {
	case "convert":
		cmds, err := batcher.GetConvertBatch()
		for _, c := range cmds {
			fmt.Println("cmd: ", c)
		}

		if err != nil || len(cmds) == 0 {
			return nil, err
		}
	case "add_fonts":
		cmds, err := batcher.GetAddFontsBatch()
		if err != nil {
			return nil, err
		}
		task.cmds = cmds

	case "add_subtitle":
		cmds, err := batcher.GetAddSubtittleBatch()
		if err != nil {
			return nil, err
		}
		task.cmds = cmds
	default:
	}

	return task, nil
}

func (t *ffmpegTask) GetId() string {
	return t.id
}

func (t *ffmpegTask) GetOB() rxgo.Observable {
	return t.ob.ob
}

func (t *ffmpegTask) GetCmds() [][]string {
	return t.cmds
}

func (t *ffmpegTask) Run() {
	if !t.exec {
		return
	}

	if err := t.batcher.ExecuteBatch(t.ob, t.ob, t.cmds); err != nil {
		t.ob.ch <- rxgo.Of("error executing batch")
		close(t.ob.ch)
		return
	}
}
