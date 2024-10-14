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
	AnimeHandler AnimeHandler
}

var HandlerSet = wire.NewSet(
	NewPingHandler,
	NewMovieHandler,
	NewStarHandler,
	NewImageHandler,
	NewTagHandler,
	NewAnimeHandler,
	wire.Struct(new(Handler), "*"),
)
