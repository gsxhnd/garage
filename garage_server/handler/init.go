package handler

import (
	"github.com/google/wire"
)

type Handler struct {
	PingHandler  PingHandler
	MovieHandler MovieHandler
	StarHandler  StarHandler
}

var HandlerSet = wire.NewSet(
	NewPingHandler,
	NewMovieHandler,
	NewStarHandler,
	wire.Struct(new(Handler), "*"),
)
