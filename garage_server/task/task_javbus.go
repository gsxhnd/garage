package task

import "github.com/reactivex/rxgo/v2"

type javbusTask struct {
	id string
	ob rxgo.Observable
	ch chan rxgo.Item
}

func NewJavbusTask() Task {
	ch := make(chan rxgo.Item)
	return &javbusTask{
		ob: rxgo.FromChannel(ch),
		ch: ch,
	}
}

func (t *javbusTask) Observable() rxgo.Observable {
	return t.ob
}

func (t *javbusTask) Run(cmd string) {}

func (t *javbusTask) GetId() string {
	return t.id
}
