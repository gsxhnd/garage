// /home/gsxhnd/Code/personal/garage/garage_server/service/star.go
package service

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/utils"
)

type StarService interface {
	CreateStars([]model.Star) error
	DeleteStars([]uint) error
	UpdateStar(model.Star) error
	GetStars(*database.Pagination) ([]model.Star, error)
}

type starService struct {
	logger utils.Logger
	db     database.Driver
}

func NewStarService(l utils.Logger, db database.Driver) StarService {
	return starService{
		logger: l,
		db:     db,
	}
}

// CreateStars implements StarService.
func (s starService) CreateStars(stars []model.Star) error {
	return s.db.CreateStars(stars)
}

// DeleteStars implements StarService.
func (s starService) DeleteStars(ids []uint) error {
	return s.db.DeleteStars(ids)
}

// UpdateStar implements StarService.
func (s starService) UpdateStar(star model.Star) error {
	panic("unimplemented")
}

// GetStars implements StarService.
func (s starService) GetStars(p *database.Pagination) ([]model.Star, error) {
	return s.db.GetStars()
}
