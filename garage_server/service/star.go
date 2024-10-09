package service

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/utils"
)

type StarService interface {
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
