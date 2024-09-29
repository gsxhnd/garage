package service

import (
	"github.com/gsxhnd/garage/garage_server/db"
	"github.com/gsxhnd/garage/utils"
)

type PingService interface {
	Ping() error
}

type pingService struct {
	logger utils.Logger
	db     *db.Database
}

func NewPingService(l utils.Logger, db *db.Database) PingService {
	return &pingService{
		logger: l,
		db:     db,
	}
}

func (p *pingService) Ping() error {
	if err := p.db.TenhouDB.Ping(); err != nil {
		p.logger.Errorf(err.Error())
		return err
	}

	return nil
}
