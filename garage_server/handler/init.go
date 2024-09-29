package handler

import (
	"github.com/google/wire"
)

type Handler struct {
	PingHandler PingHandler
}

var HandlerSet = wire.NewSet(
	NewPingHandler,
	wire.Struct(new(Handler), "*"),
)
