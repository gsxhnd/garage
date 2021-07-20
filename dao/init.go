package dao

import (
	"garage/db"
	"github.com/gsxhnd/owl"
	"gorm.io/gorm"
)

var Database datebase

type datebase struct {
	Default *gorm.DB
}

func (d *datebase) ConnectPostgreSQL() error {
	var err error
	d.Default, err = db.GetPostgreSQL()
	if err != nil {
		return err
	}
	return nil
}

func (d *datebase) ConnectMySQL() error {
	var err error
	d.Default, err = db.GetMariaDB(
		owl.GetString("db.mysql.user"),
		owl.GetString("db.mysql.password"),
		owl.GetString("db.mysql.addr"),
		owl.GetString("db.mysql.dbname"))
	if err != nil {
		return err
	}
	return nil
}

func (d *datebase) Close() {
	m, _ := d.Default.DB()
	_ = m.Close()
}
