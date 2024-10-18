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
	UpdateActor(actor *model.Actor) error
	GetActors() ([]model.Actor, error)
	SearchActorByName(string) ([]model.Actor, error)
	CreateTags([]model.Tag) error
	DeleteTags([]uint) error
	UpdateTag(tag *model.Tag) error
	GetTags() ([]model.Tag, error)
	SearchTagsByName(name string) ([]model.Tag, error)
	CreateMovieActors(movieActors []model.MovieActor) error
	DeleteMovieActors(ids []uint) error
	UpdateMovieActor(model.MovieActor) error
	GetMovieActors() ([]model.MovieActor, error)
	GetMovieActorsByMovieId(id uint) ([]model.MovieActor, error)
	CreateMovieTags([]model.MovieTag) error
	DeleteMovieTags(ids []uint) error
	GetMovieTagByMovieId(movieId uint) ([]model.MovieTag, error)
	UpdateMovieTag(model.MovieTag) error
	CreateAnimes([]model.Anime) error
	DeleteAnimes([]uint) error
	UpdateAnime(model.Anime) error
	GetAnimes(*Pagination) ([]model.Anime, error)
}
