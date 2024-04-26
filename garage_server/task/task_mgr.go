package task

import "github.com/gsxhnd/garage/utils"

type TaskMgr interface {
	GetTaskList(string)
	AddTask(Task)
}

type taskManager struct {
	logger utils.Logger
	tasks  map[string]Task
}

func NewTaskMgr(l utils.Logger) TaskMgr {
	return &taskManager{
		logger: l,
		tasks:  make(map[string]Task, 0),
	}
}

func (m *taskManager) GetTaskList(id string) {
	m.logger.Debugf("GetTaskList id: %s", id)
}

func (m *taskManager) AddTask(task Task) {
	m.tasks[task.GetId()] = task
	go task.Run()
}
