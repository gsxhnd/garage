package service

import "github.com/google/wire"

var ServiceSet = wire.NewSet(
	NewPingService,
	NewMovieService,
	NewMovieActorService,
	NewMovieTagService,
	NewActorService,
	NewTagService,
	NewAnimeService,
)
