package dao

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
	"github.com/google/wire"
	"github.com/gsxhnd/garage/utils"
)

type Database struct {
	sqliteDB *sql.DB
}

func NewDatabase(cfg *utils.Config, l utils.Logger) (*Database, error) {
	sqliteDB, err := sql.Open("sqlite", "./data/billfish.db")
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		sqliteDB: sqliteDB,
	}, nil
}

var DaoSet = wire.NewSet(NewDatabase, NewTestDao)
