package dao

import "garage/model"

func GetAllJavTag() ([]model.JavTag, error) {
	var (
		db   = Database.Default
		data []model.JavTag
	)
	row := db.Model(&model.JavTag{}).Find(&data)
	if row.Error != nil {
		return nil, row.Error
	} else {
		return data, nil
	}
}
