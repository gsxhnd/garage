package handler

import (
	"github.com/google/wire"
)

type Handler struct {
	RootHandler      RootHandler
	WebsocketHandler WebsocketHandler
	JavHandler       JavHandler
}

var HandlerSet = wire.NewSet(
	NewPingHandle,
	NewWebsocketHandler,
	NewJavHandler,
	wire.Struct(new(Handler), "*"),
)
