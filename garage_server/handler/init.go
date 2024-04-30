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
	NewRootHandler,
	NewWebsocketHandler,
	NewJavHandler,
	NewFFmpegHandler,
	NewCrawlHandler,
	wire.Struct(new(Handler), "*"),
)
