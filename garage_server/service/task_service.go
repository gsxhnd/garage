package service

type TaskService interface{}

type taskService struct{}

func NewTaskService() TaskService {
	return &taskService{}
}
