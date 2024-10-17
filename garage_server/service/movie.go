package service

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/utils"
)

type MovieService interface {
	CreateMovies([]model.Movie) error
	DeleteMovies([]uint) error
	UpdateMovie(model.Movie) error
	GetMovies(*database.Pagination) ([]model.Movie, error)
	GetMovieInfo(code string) (*model.MovieInfo, error)
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
func (s movieService) CreateMovies(movies []model.Movie) error {
	return s.db.CreateMovies(movies)
}

// DeleteMovies implements MovieService.
func (s movieService) DeleteMovies(ids []uint) error {
	return s.db.DeleteMovies(ids)
}

// UpdateMovie implements MovieService.
func (s movieService) UpdateMovie(model.Movie) error {
	panic("unimplemented")
}

// GetMovies implements MovieService.
func (s movieService) GetMovies(p *database.Pagination) ([]model.Movie, error) {
	return s.db.GetMovies(p)
}

func (s movieService) GetMovieInfo(code string) (*model.MovieInfo, error) {
	var data model.MovieInfo
	movie, err := s.db.GetMovieByCode(code)
	if err != nil {
		return nil, err
	}
	data.Movie = *movie

	return &data, nil
}
