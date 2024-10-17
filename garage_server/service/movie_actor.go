// garage_server/service/movie_actor.go
package service

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/utils"
)

type MovieActorService interface {
	CreateMovieActors(data []model.MovieActor) error
	DeleteMovieActors(ids []uint) error
	UpdateMovieActor(model.MovieActor) error
	GetMovieActors() ([]model.MovieActor, error)
}

type movieActorService struct {
	logger utils.Logger
	db     database.Driver
}

func NewMovieActorService(l utils.Logger, db database.Driver) MovieActorService {
	return movieActorService{
		logger: l,
		db:     db,
	}
}

// CreateMovieActors implements MovieActorService.
func (s movieActorService) CreateMovieActors(movieActors []model.MovieActor) error {
	return s.db.CreateMovieActors(movieActors)
}

// DeleteMovieActors implements MovieActorService.
func (s movieActorService) DeleteMovieActors(ids []uint) error {
	return s.db.DeleteMovieActors(ids)
}

// UpdateMovieActor implements MovieActorService.
func (s movieActorService) UpdateMovieActor(movieActor model.MovieActor) error {
	panic("unimplemented")
}

// GetMovieActors implements MovieActorService.
func (s movieActorService) GetMovieActors() ([]model.MovieActor, error) {
	return s.db.GetMovieActors()
}
