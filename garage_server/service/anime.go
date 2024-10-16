package service

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/utils"
)

type AnimeService interface {
	CreateAnimes([]model.Anime) error
	DeleteAnimes([]uint) error
	UpdateAnime(model.Anime) error
	GetAnimes(*database.Pagination) ([]model.Anime, error)
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
	return s.db.CreateAnimes(data)
}

// DeleteAnimes implements AnimeService.
func (s animeService) DeleteAnimes(ids []uint) error {
	return s.db.DeleteAnimes(ids)
}

// UpdateAnime implements AnimeService.
func (s animeService) UpdateAnime(data model.Anime) error {
	return s.db.UpdateAnime(data)
}

// GetAnimes implements AnimeService.
func (s animeService) GetAnimes(p *database.Pagination) ([]model.Anime, error) {
	return s.db.GetAnimes(p)
}
