package service

import (
	"github.com/gsxhnd/garage/garage_server/db"
	"github.com/gsxhnd/garage/garage_server/storage"
	"github.com/gsxhnd/garage/utils"
)

type PingService interface {
	Ping() error
}

type pingService struct {
	logger  utils.Logger
	db      *db.Database
	storage storage.Storage
}

func NewPingService(l utils.Logger, db *db.Database, s storage.Storage) PingService {
	return &pingService{
		logger:  l,
		db:      db,
		storage: s,
	}
}

func (p *pingService) Ping() error {
	if err := p.db.SqliteDB.Ping(); err != nil {
		p.logger.Errorf(err.Error())
		return err
	}

	if err := p.storage.Ping(); err != nil {
		p.logger.Errorf(err.Error())
		return err
	}

	return nil
}
