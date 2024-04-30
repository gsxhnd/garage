package handler

import (
	"github.com/google/wire"
)

type Handler struct {
	RootHandler      RootHandler
	WebsocketHandler WebsocketHandler
	CrawlHandler     CrawlHandler
	JavHandler       JavHandler
	FFmpegHander     FFmpegHandler
}

var HandlerSet = wire.NewSet(
	NewRootHandle,
	NewWebsocketHandler,
	NewJavHandler,
	NewFFmpegHandler,
	NewCrwalHandler,
	wire.Struct(new(Handler), "*"),
)
