package dao

import "garage/model"

func CreateJavStar(js *[]model.JavStar) error {
	var db = Database.Default
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

func UpdateJavStar(data model.JavStar) error {
	var db = Database.Default
	row := db.Model(model.JavStar{}).Where("id = ?", data.Id).Updates(model.JavStar{
		Name:               data.Name,
		JavbusStarCode:     data.JavbusStarCode,
		JavlibraryStarCode: data.JavlibraryStarCode,
		Image:              data.Image,
	})
	if row.Error != nil {
		return row.Error
	}
	return nil
}

func DelJavStar(id uint64) error {
	var db = Database.Default
	row := db.Where("id = ?", id).Delete(&model.JavStar{})
	if row.Error != nil {
		return row.Error
	}
	return nil
}
