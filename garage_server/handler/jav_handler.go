package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/task"
	"github.com/gsxhnd/garage/utils"
)

type JavHandler interface {
	CrawlJavByCode(ctx *fiber.Ctx) error
}

type javHnadler struct {
	logger    utils.Logger
	validator *validator.Validate
	taskMgr   task.TaskMgr
}

func NewJavHandler(l utils.Logger, v *validator.Validate, t task.TaskMgr) JavHandler {
	return &javHnadler{
		logger:    l,
		validator: v,
		taskMgr:   t,
	}
}

func (h *javHnadler) CrawlJavByCode(ctx *fiber.Ctx) error {
	return ctx.SendString("CrawlJavByCode")
}
