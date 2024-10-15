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
	SearchStarByName(name string) ([]model.Star, error)
	CreateTags(tags []model.Tag) error
	DeleteTags(ids []uint) error
	GetTags() ([]model.Tag, error)
	CreateMovieStars(movieStars []model.MovieStar) error
	DeleteMovieStars(ids []uint) error
	GetMovieStars() ([]model.MovieStar, error)
	UpdateMovieStar(movieStar model.MovieStar) error
	CreateMovieTags(movieTags []model.MovieTag) error
	DeleteMovieTags(ids []uint) error
	GetMovieTags() ([]model.MovieTag, error)
	UpdateMovieTag(movieTag model.MovieTag) error
}

type Pagination struct {
	Limit  uint
	Offset uint
}
