package service

import (
	"github.com/gsxhnd/garage/garage_server/db"
	"github.com/gsxhnd/garage/utils"
)

type TaskService interface{}

type taskService struct {
	logger utils.Logger
	db     db.TaskDao
}

func NewTaskService(l utils.Logger, db db.TaskDao) TaskService {
	return &taskService{
		logger: l,
		db:     db,
	}
}
