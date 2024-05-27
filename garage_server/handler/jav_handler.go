package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/garage_server/task"
	"github.com/gsxhnd/garage/utils"
)

type JavHandler interface {
	CrawlJavbus(ctx *fiber.Ctx) error
}

type javHnadler struct {
	logger    utils.Logger
	validator *validator.Validate
	taskMgr   task.TaskMgr
	svc       service.TaskService
}

func NewJavHandler(l utils.Logger, v *validator.Validate, t task.TaskMgr, svc service.TaskService) JavHandler {
	return &javHnadler{
		logger:    l,
		validator: v,
		taskMgr:   t,
		svc:       svc,
	}
}

func (h *javHnadler) CrawlJavbus(ctx *fiber.Ctx) error {
	return nil
}
