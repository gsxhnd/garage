package dao

import (
	"garage/db"
	"garage/model"
	"gorm.io/gorm"
	"log"
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
