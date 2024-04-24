package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/utils"
)

type RootHandler interface {
	Ping(ctx *fiber.Ctx) error
}

type rootHandle struct {
	logger utils.Logger
	svc    service.TestService
}

func NewPingHandle(l utils.Logger, s service.TestService) RootHandler {
	return &rootHandle{
		logger: l,
		svc:    s,
	}
}

func (h *rootHandle) Ping(ctx *fiber.Ctx) error {
	// ctx.JSON(200, "pong")
	ctx.Status(200).JSON(fiber.Map{
		
	})
	return ctx.SendString("pong")
}
