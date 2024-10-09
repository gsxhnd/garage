package service

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/utils"
)

type MovieService interface {
}

type movieService struct {
	logger utils.Logger
	db     database.Driver
}

func NewMovieService(l utils.Logger, db database.Driver) MovieService {
	return movieService{
		logger: l,
		db:     db,
	}
}

func (s *movieService) AddMovies() error {
	return nil
}
