package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/garage/utils"
)

type Middlewarer interface {
	RequestLog() gin.HandlerFunc
}
type middleware struct {
	logger utils.Logger
}

func NewMiddleware(l utils.Logger) Middlewarer {
	return &middleware{
		logger: l,
	}
}
