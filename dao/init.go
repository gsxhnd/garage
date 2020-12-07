package dao

import (
	"garage/db"
	"gorm.io/gorm"
)

var Database datebase

type datebase struct {
	Default *gorm.DB
}

func (d *datebase) Connect() error {
	var err error
	d.Default, err = db.GetSqliteDB()
	if err != nil {
		return err
	}
	return nil
}

func (d *datebase) Close() {
	m, _ := d.Default.DB()
	_ = m.Close()
}
