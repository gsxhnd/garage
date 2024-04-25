package task

import "github.com/google/wire"

type Task interface {
	GetId() string
	Run(cmd string)
}

var TaskSet = wire.NewSet(NewTaskMgr)
