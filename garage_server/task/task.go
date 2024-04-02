package task

type Task interface{}

type TaskMgr interface {
	GetTaskList()
}

type taskManager struct {
	tasks map[string]Task
}

func NewTaskMgr() TaskMgr {
	return &taskManager{
		tasks: make(map[string]Task, 0),
	}
}

func (m *taskManager) GetTaskList() {}
