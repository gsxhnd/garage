package middleware

import (
	"github.com/gsxhnd/garage/utils"
)

type Middlewarer interface {
	// RequestLog(ctx *fiber.Ctx) error
}
type middleware struct {
	logger utils.Logger
}

func NewMiddleware(l utils.Logger) Middlewarer {
	return &middleware{
		logger: l,
	}
}
