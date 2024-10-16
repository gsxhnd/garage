package database

import "github.com/gsxhnd/garage/garage_server/model"

type Driver interface {
	Ping() error
	Migrate() error
	CreateMovies([]model.Movie) error
	DeleteMovies([]uint) error
	GetMovies(*Pagination) ([]model.Movie, error)
	CreateStars([]model.Star) error
	DeleteStars([]uint) error
	GetStars() ([]model.Star, error)
	SearchStarByName(string) ([]model.Star, error)
	CreateTags([]model.Tag) error
	DeleteTags([]uint) error
	GetTags() ([]model.Tag, error)
	CreateMovieStars(movieStars []model.MovieStar) error
	DeleteMovieStars(ids []uint) error
	GetMovieStars() ([]model.MovieStar, error)
	UpdateMovieStar(model.MovieStar) error
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
	Limit  uint
	Offset uint
}
