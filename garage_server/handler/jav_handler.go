package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/utils"
)

type JavHandler interface {
	CrawlJavByCode(ctx *fiber.Ctx) error
}

type javHnadler struct {
	logger utils.Logger
}

func NewJavHandler(l utils.Logger) JavHandler {
	return &javHnadler{
		logger: l,
	}
}

func (h *javHnadler) CrawlJavByCode(ctx *fiber.Ctx) error {
	return ctx.SendString("CrawlJavByCode")
}
