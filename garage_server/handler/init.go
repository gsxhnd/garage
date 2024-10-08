package handler

import (
	"github.com/google/wire"
)

type Handler struct {
	PingHandler  PingHandler
	MovieHandler MovieHandler
	StarHandler  StarHandler
	ImageHandler ImageHandler
}

var HandlerSet = wire.NewSet(
	NewPingHandler,
	NewMovieHandler,
	NewStarHandler,
	NewImageHandler,
	wire.Struct(new(Handler), "*"),
)
