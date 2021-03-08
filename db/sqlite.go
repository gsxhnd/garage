package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetSqliteDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("jav.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
