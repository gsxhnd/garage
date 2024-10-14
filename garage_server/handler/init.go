package handler

import (
	"github.com/google/wire"
)

type Handler struct {
	PingHandler  PingHandler
	MovieHandler MovieHandler
	StarHandler  StarHandler
	ImageHandler ImageHandler
	TagHandler   TagHandler
}

var HandlerSet = wire.NewSet(
	NewPingHandler,
	NewMovieHandler,
	NewStarHandler,
	NewImageHandler,
	NewTagHandler,
	wire.Struct(new(Handler), "*"),
)
