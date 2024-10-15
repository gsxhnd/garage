package service

import "github.com/google/wire"

var ServiceSet = wire.NewSet(
	NewPingService,
	NewMovieService,
	NewMovieStarService,
	NewMovieTagService,
	NewStarService,
	NewTagService,
	NewAnimeService,
)
