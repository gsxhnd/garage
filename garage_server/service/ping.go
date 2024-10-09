package service

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/storage"
	"github.com/gsxhnd/garage/utils"
)

type PingService interface {
	Ping() error
}

type pingService struct {
	logger  utils.Logger
	db      database.Driver
	storage storage.Storage
}

func NewPingService(l utils.Logger, db database.Driver, s storage.Storage) PingService {
	return &pingService{
		logger:  l,
		db:      db,
		storage: s,
	}
}

func (p *pingService) Ping() error {
	if err := p.db.Ping(); err != nil {
		p.logger.Errorf(err.Error())
		return err
	}

	if err := p.storage.Ping(); err != nil {
		p.logger.Errorf(err.Error())
		return err
	}

	return nil
}
