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

func UpdateJavMovie(m model.JavMovie) error {
	var (
		db = Database.Default
	)
	row := db.Model(&model.JavMovie{}).Where("code = ?", m.Code).Updates(m)
	return row.Error
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

func GetJavMovieInfo(code string) (interface{}, error) {
	var (
		db   = Database.Default
		data = struct {
			model.JavMovie
			Star []model.JavStar `json:"star"`
			Tag  []model.JavTag  `json:"tag"`
		}{}
	)
	if row := db.Model(&model.JavMovie{}).
		Where("code = ?", code).
		Find(&data); row.Error != nil {
		return nil, row.Error
	}

	if row := db.Model(&model.JavMovieSatr{}).
		Where("code = ?", code).
		Find(&data.Star); row.Error != nil {
		return nil, row.Error
	}

	if row := db.Model(&model.JavMovieTag{}).
		Where("jav_movie_code = ?", code).
		Find(&data.Tag); row.Error != nil {
		return nil, row.Error
	}
	return data, nil
}

func DelJavMovie(code string) error {
	var db = Database.Default
	row := db.Where("code = ?", code).Delete(&model.JavMovie{})
	return row.Error
}
