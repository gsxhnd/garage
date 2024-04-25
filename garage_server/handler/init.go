package handler

import (
	"github.com/google/wire"
)

type Handler struct {
	RootHandler      RootHandler
	WebsocketHandler WebsocketHandler
	JavHandler       JavHandler
	FFmpegHander     FFmpegHandler
}

var HandlerSet = wire.NewSet(
	NewPingHandle,
	NewWebsocketHandler,
	NewJavHandler,
	NewFFmpegHandler,
	wire.Struct(new(Handler), "*"),
)
