package db

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/gsxhnd/garage/utils"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	SqliteDB *sql.DB
	logger   utils.Logger
}

func NewDatabase(cfg *utils.Config, l utils.Logger) (*Database, error) {
	db, err := sql.Open("sqlite3", cfg.DatabasePath)
	if err != nil {
		return nil, err
	}

	return &Database{
		SqliteDB: db,
		logger:   l,
	}, nil
}

var DBSet = wire.NewSet(NewDatabase)
