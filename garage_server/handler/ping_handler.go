package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/utils"
)

type RootHandler interface {
	Ping(ctx *fiber.Ctx) error
}

type rootHandle struct {
	logger utils.Logger
}

func NewPingHandle(l utils.Logger) RootHandler {
	return &rootHandle{
		logger: l,
	}
}

func (h *rootHandle) Ping(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}
