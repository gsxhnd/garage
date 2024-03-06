package garage_jav

import (
	"github.com/reactivex/rxgo/v2"
)

type Task struct {
	id   string
	name string
	data []JavMovie
	ob   rxgo.Observable
}

func NewTask() *Task {
	return &Task{
		id:   "",
		name: "",
		data: make([]JavMovie, 0),
		ob:   rxgo.Empty(),
	}
}

func (t *Task) GetTaskInfo() {
}
