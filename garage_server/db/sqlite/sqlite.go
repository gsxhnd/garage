package sqlite

import (
	"database/sql"

	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/utils"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteDB struct {
	db     *sql.DB
	logger utils.Logger
}

func NewSqliteDB(dataSource string, l utils.Logger) (database.Driver, error) {
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		return nil, err
	}

	return &sqliteDB{
		db:     db,
		logger: l,
	}, nil
}

func (db *sqliteDB) Ping() error {
	return db.db.Ping()
}
