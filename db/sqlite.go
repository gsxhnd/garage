package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetSqliteDB(file string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
