package garage_task

import "github.com/reactivex/rxgo/v2"

type Task interface{}

type task struct {
	ob rxgo.Observable
	ch chan rxgo.Item
}

func NewTask() Task {
	ch := make(chan rxgo.Item)
	return &task{
		ob: rxgo.FromChannel(ch),
		ch: ch,
	}
}

func (t *task) Observable() rxgo.Observable {
	return t.ob
}

func (t *task) RunFFmpeg() {}

func (t *task) RunJavbusCrawl() {}
