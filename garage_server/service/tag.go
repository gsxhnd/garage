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
	UpdateTag(*model.Tag) error
	GetTags() ([]model.Tag, error)
	SearchTagsByName(name string) ([]model.Tag, error)
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
func (s tagService) CreateTags(tags []model.Tag) error {
	return s.db.CreateTags(tags)
}

// DeleteTags implements TagService.
func (s tagService) DeleteTags(ids []uint) error {
	return s.db.DeleteTags(ids)
}

// UpdateTag implements TagService.
func (s tagService) UpdateTag(tag *model.Tag) error {
	return s.db.UpdateTag(tag)
}

// GetTags implements TagService.
func (s tagService) GetTags() ([]model.Tag, error) {
	return s.db.GetTags()
}

func (s tagService) SearchTagsByName(name string) ([]model.Tag, error) {
	return s.db.SearchTagsByName(name)
}
