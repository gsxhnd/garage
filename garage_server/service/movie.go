package service

import (
	"github.com/gsxhnd/garage/garage_server/db"
	"github.com/gsxhnd/garage/utils"
)

type MovieService interface {
	Ping() error
}

type movieService struct {
	logger utils.Logger
	db     *db.Database
}

func NewMovieService(l utils.Logger, db *db.Database) MovieService {
	return movieService{
		logger: l,
		db:     db,
	}
}

func (p movieService) Ping() error {
	if err := p.db.TenhouDB.Ping(); err != nil {
		p.logger.Errorf(err.Error())
		return err
	}

	return nil
}
