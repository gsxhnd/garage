// garage_server/service/movie_star.go
package service

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/utils"
)

type MovieStarService interface {
	CreateMovieStars(data []model.MovieStar) error
	DeleteMovieStars(ids []uint) error
	UpdateMovieStar(model.MovieStar) error
	GetMovieStars() ([]model.MovieStar, error)
}

type movieStarService struct {
	logger utils.Logger
	db     database.Driver
}

func NewMovieStarService(l utils.Logger, db database.Driver) MovieStarService {
	return movieStarService{
		logger: l,
		db:     db,
	}
}

// CreateMovieStars implements MovieStarService.
func (s movieStarService) CreateMovieStars(movieStars []model.MovieStar) error {
	return s.db.CreateMovieStars(movieStars)
}

// DeleteMovieStars implements MovieStarService.
func (s movieStarService) DeleteMovieStars(ids []uint) error {
	return s.db.DeleteMovieStars(ids)
}

// UpdateMovieStar implements MovieStarService.
func (s movieStarService) UpdateMovieStar(movieStar model.MovieStar) error {
	panic("unimplemented")
}

// GetMovieStars implements MovieStarService.
func (s movieStarService) GetMovieStars() ([]model.MovieStar, error) {
	return s.db.GetMovieStars()
}
