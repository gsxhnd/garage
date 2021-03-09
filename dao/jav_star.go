package dao

import "garage/model"

func CreateJavStar(js *[]model.JavStar) error {
	db := Database.Default
	row := db.Create(&js)
	if row.Error != nil {
		return nil
	} else {
		return row.Error
	}
}

func GetJavStarList() (interface{}, error) {
	var (
		db   = Database.Default
		data = struct {
			Total int64           `json:"total"`
			List  []model.JavStar `json:"list"`
		}{}
	)
	row := db.Model(&model.JavStar{}).
		Count(&data.Total).
		Find(&data.List)
	if row.Error != nil {
		return nil, row.Error
	} else {
		return data, nil
	}
}
