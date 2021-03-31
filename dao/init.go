package dao

import (
	"garage/db"
	"garage/model"
	"github.com/gsxhnd/owl"
	"gorm.io/gorm"
	"log"
)

var Database datebase

type datebase struct {
	Default *gorm.DB
}

func (d *datebase) ConnectSQLite(file string) error {
	var err error
	d.Default, err = db.GetSqliteDB(file)
	if err != nil {
		return err
	}

	err = d.Default.AutoMigrate(&model.JavMovie{}, &model.JavStar{}, &model.JavMovieSatr{})
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (d *datebase) ConnectPostgreSQL() error {
	var err error
	d.Default, err = db.GetPostgreSQL()
	if err != nil {
		return err
	}

	err = d.Default.AutoMigrate(&model.JavMovie{}, &model.JavStar{}, &model.JavMovieSatr{})
	if err != nil {
		log.Println(err)
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

	err = d.Default.AutoMigrate(&model.JavMovie{}, &model.JavStar{}, &model.JavMovieSatr{})
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (d *datebase) Close() {
	m, _ := d.Default.DB()
	_ = m.Close()
}
