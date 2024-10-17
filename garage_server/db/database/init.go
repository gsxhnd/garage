package database

import "github.com/gsxhnd/garage/garage_server/model"

type Driver interface {
	Ping() error
	Migrate() error
	CreateMovies([]model.Movie) error
	DeleteMovies([]uint) error
	GetMovies(*Pagination) ([]model.Movie, error)
	GetMovieByCode(code string) (*model.Movie, error)
	CreateActors([]model.Actor) error
	DeleteActors([]uint) error
	GetActors() ([]model.Actor, error)
	SearchActorByName(string) ([]model.Actor, error)
	CreateTags([]model.Tag) error
	DeleteTags([]uint) error
	GetTags() ([]model.Tag, error)
	CreateMovieActors(movieActors []model.MovieActor) error
	DeleteMovieActors(ids []uint) error
	GetMovieActors() ([]model.MovieActor, error)
	UpdateMovieActor(model.MovieActor) error
	CreateMovieTags([]model.MovieTag) error
	DeleteMovieTags(ids []uint) error
	GetMovieTags() ([]model.MovieTag, error)
	UpdateMovieTag(model.MovieTag) error
	CreateAnimes([]model.Anime) error
	DeleteAnimes([]uint) error
	UpdateAnime(model.Anime) error
	GetAnimes(*Pagination) ([]model.Anime, error)
}

type Pagination struct {
	Limit  uint `validate:"max=100,min=1,number"`
	Offset uint `validate:"min=0,number"`
}
