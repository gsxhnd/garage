package database

import "github.com/gsxhnd/garage/garage_server/model"

type Driver interface {
	Ping() error
	GetMovies() ([]model.Movie, error)
	CreateMovies(movies []model.Movie) error
	Migrate() error
}
