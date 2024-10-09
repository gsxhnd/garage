package db

import (
	"github.com/google/wire"
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/db/sqlite"
	"github.com/gsxhnd/garage/utils"
)

func NewDatabase(cfg *utils.Config, l utils.Logger) (database.Driver, error) {
	d, err := sqlite.NewSqliteDB(cfg.DatabasePath, l)
	if err != nil {
		return nil, err
	}
	if err := d.Migrate(); err != nil {
		return nil, err
	}
	return d, nil
}

var DBSet = wire.NewSet(NewDatabase)
