package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/utils"
)

type PingHandler interface {
	Ping(ctx *gin.Context)
}
type rootHandle struct {
	logger utils.Logger
	svc    service.TestService
}

func NewPingHandle(l utils.Logger, s service.TestService) PingHandler {
	return &rootHandle{
		logger: l,
		svc:    s,
	}
}

func (h *rootHandle) Ping(ctx *gin.Context) {
	ctx.JSON(200, "pong")
}
