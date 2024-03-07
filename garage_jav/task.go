package garage_jav

import (
	"github.com/gocolly/colly/v2"
	"github.com/reactivex/rxgo/v2"
)

type Task struct {
	id        string
	name      string
	data      []JavMovie
	collector *colly.Collector
	ob        rxgo.Observable
}

func NewTask(collector *colly.Collector) *Task {
	return &Task{
		id:        "",
		name:      "",
		data:      make([]JavMovie, 0),
		collector: collector,
		ob:        rxgo.Empty(),
	}
}

func (t *Task) GetTaskInfo() {
}
