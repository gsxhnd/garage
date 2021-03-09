package dao

import (
	"garage/model"
)

func CreateJavMovie(jm *[]model.JavMovie) error {
	db := Database.Default
	row := db.Create(&jm)
	if row.Error != nil {
		return nil
	} else {
		return row.Error
	}
}

func GetJavMovie() (interface{}, error) {
	var (
		db   = Database.Default
		data = struct {
			Total int64            `json:"total"`
			List  []model.JavMovie `json:"list"`
		}{}
	)
	row := db.Model(&model.JavMovie{}).
		Count(&data.Total).
		Find(&data.List)
	if row.Error != nil {
		return nil, row.Error
	} else {
		return data, nil
	}
}