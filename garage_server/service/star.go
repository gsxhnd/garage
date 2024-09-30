package service

import (
	"github.com/gsxhnd/garage/garage_server/db"
	"github.com/gsxhnd/garage/utils"
)

type StarService interface {
}

type starService struct {
	logger utils.Logger
	db     *db.Database
}

func NewStarService(l utils.Logger, db *db.Database) StarService {
	return starService{
		logger: l,
		db:     db,
	}
}
