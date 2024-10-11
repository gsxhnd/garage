package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/errno"
	"github.com/gsxhnd/garage/garage_server/service"
)

type PingHandler interface {
	Ping(ctx *fiber.Ctx) error
}

type pingHandle struct {
	svc service.PingService
}

func NewPingHandler(svc service.PingService) PingHandler {
	return &pingHandle{
		svc: svc,
	}
}

// @Description  ping serivce working, db connect
// @Produce      json
// @Success      200
// @Router       /ping [get]
func (h *pingHandle) Ping(ctx *fiber.Ctx) error {
	err := h.svc.Ping()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.JSON(errno.DecodeError(errno.UnknownError))
}
