package dao

import (
	"fmt"
	"garage/model"
)

func GetJavMovie() interface{} {
	var (
		db     = Database.Default
		movies []model.JavMovie
	)

	row := db.Model(&model.JavMovie{}).Find(&movies)
	if row.Error != nil {
		return nil
	} else {
		fmt.Println(movies)
		return movies
	}
}
