package sqlite

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	migrate_sqlite "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
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

//go:embed migrations/*.sql
var s embed.FS

func (db *sqliteDB) Migrate() error {
	d, err := iofs.New(s, "migrations")
	if err != nil {
		db.logger.Errorw("", "error", err)
		return err
	}

	driver, err := migrate_sqlite.WithInstance(db.db, &migrate_sqlite.Config{})
	if err != nil {
		db.logger.Errorw("", "error", err)
		return err
	}
	m, err := migrate.NewWithInstance("iofs", d, "sqlite3", driver)
	if err != nil {
		db.logger.Errorw("", "error", err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
