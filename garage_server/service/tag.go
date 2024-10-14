// garage_server/service/tag.go
package service

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/utils"
)

type TagService interface {
	CreateTags(data []model.Tag) error
	DeleteTags(ids []uint) error
	UpdateTag(model.Tag) error
	GetTags() ([]model.Tag, error)
}

type tagService struct {
	logger utils.Logger
	db     database.Driver
}

func NewTagService(l utils.Logger, db database.Driver) TagService {
	return tagService{
		logger: l,
		db:     db,
	}
}

// CreateTags implements TagService.
func (s tagService) CreateTags(data []model.Tag) error {
	panic("unimplemented")
}

// DeleteTags implements TagService.
func (s tagService) DeleteTags(ids []uint) error {
	panic("unimplemented")
}

// UpdateTag implements TagService.
func (s tagService) UpdateTag(tag model.Tag) error {
	panic("unimplemented")
}

// GetTags implements TagService.
func (s tagService) GetTags() ([]model.Tag, error) {
	panic("unimplemented")
}
