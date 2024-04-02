package handler

import (
	"github.com/google/wire"
	"github.com/gsxhnd/garage/utils"
)

type Handler struct {
	PingHandler      PingHandler
	WebsocketHandler WebsocketHandler
}

func NewHandler(cfg *utils.Config) *Handler {
	return &Handler{}
}

// var HandlerSet = wire.NewSet(
// 	NewPingHandle,
// )

var HandlerSet = wire.NewSet(
	NewPingHandle,
	NewWebsocketHandler,
	wire.Struct(new(Handler), "*"),
)
