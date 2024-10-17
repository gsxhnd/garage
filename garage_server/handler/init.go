package handler

import (
	"github.com/google/wire"
)

type Handler struct {
	PingHandler      PingHandler
	MovieHandler     MovieHandler
	MovieActorHandle MovieActorHandler
	MovieTagHandler  MovieTagHandler
	ActorHandler     ActorHandler
	ImageHandler     ImageHandler
	TagHandler       TagHandler
	AnimeHandler     AnimeHandler
}

var HandlerSet = wire.NewSet(
	NewPingHandler,
	NewMovieHandler,
	NewMovieActorHandler,
	NewMovieTagHandler,
	NewActorHandler,
	NewImageHandler,
	NewTagHandler,
	NewAnimeHandler,
	wire.Struct(new(Handler), "*"),
)
