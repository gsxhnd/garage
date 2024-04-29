package middleware

import "github.com/gofiber/fiber/v2"

func (m *middleware) Websocket(ctx *fiber.Ctx) error {
	ctx.Next()
	return nil
}
