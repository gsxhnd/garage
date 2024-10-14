package database

import "github.com/gsxhnd/garage/garage_server/model"

type Driver interface {
	Ping() error
	Migrate() error
	CreateMovies(movies []model.Movie) error
	DeleteMovies(ids []uint) error
	GetMovies(*Pagination) ([]model.Movie, error)
	CreateStars(stars []model.Star) error
	DeleteStars(ids []uint) error
	GetStars() ([]model.Star, error)
	CreateTags(tags []model.Tag) error
	DeleteTags(ids []uint) error
	GetTags() ([]model.Tag, error)
}

type Pagination struct {
	Limit  uint
	Offset uint
}
