package task

import (
	"github.com/google/wire"
	"github.com/reactivex/rxgo/v2"
)

type Task interface {
	GetId() string
	Run()
	GetOB() rxgo.Observable
}

type taskOb struct {
	ob rxgo.Observable
	ch chan rxgo.Item
}

func newTaskOb() *taskOb {
	ch := make(chan rxgo.Item)
	return &taskOb{
		ob: rxgo.FromChannel(ch),
		ch: ch,
	}
}

func (o *taskOb) Write(p []byte) (int, error) {
	o.ch <- rxgo.Of(string(p))
	return len(p), nil
}

var TaskSet = wire.NewSet(NewTaskMgr)
