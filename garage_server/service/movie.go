package service

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/utils"
)

type MovieService interface {
	CreateMovies(data []model.Movie) error
	DeleteMovies(ids []uint) error
	UpdateMovie(model.Movie) error
	GetMovies() ([]model.Movie, error)
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

// CreateMovies implements MovieService.
func (s movieService) CreateMovies(data []model.Movie) error {
	panic("unimplemented")
}

// DeleteMovies implements MovieService.
func (s movieService) DeleteMovies(ids []uint) error {
	panic("unimplemented")
}

// UpdateMovie implements MovieService.
func (s movieService) UpdateMovie(model.Movie) error {
	panic("unimplemented")
}

// GetMovies implements MovieService.
func (s movieService) GetMovies() ([]model.Movie, error) {
	panic("unimplemented")
}
