package service

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/utils"
)

type AnimeService interface {
	CreateAnimes(data []model.Anime) error
	DeleteAnimes(ids []uint) error
	UpdateAnime(model.Anime) error
	GetAnimes() ([]model.Anime, error)
}

type animeService struct {
	logger utils.Logger
	db     database.Driver
}

func NewAnimeService(l utils.Logger, db database.Driver) AnimeService {
	return animeService{
		logger: l,
		db:     db,
	}
}

// CreateAnimes implements AnimeService.
func (s animeService) CreateAnimes(data []model.Anime) error {
	panic("unimplemented")
}

// DeleteAnimes implements AnimeService.
func (s animeService) DeleteAnimes(ids []uint) error {
	panic("unimplemented")
}

// UpdateAnime implements AnimeService.
func (s animeService) UpdateAnime(model.Anime) error {
	panic("unimplemented")
}

// GetAnimes implements AnimeService.
func (s animeService) GetAnimes() ([]model.Anime, error) {
	panic("unimplemented")
}
