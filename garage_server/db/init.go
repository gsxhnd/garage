package db

import (
	"github.com/google/wire"
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/db/sqlite"
	"github.com/gsxhnd/garage/utils"
)

func NewDatabase(cfg *utils.Config, l utils.Logger) (database.Driver, error) {
	return sqlite.NewSqliteDB(cfg.DatabasePath, l)
}

var DBSet = wire.NewSet(NewDatabase)
