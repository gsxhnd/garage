package dao

import (
	"garage/model"
	"log"
)

func GetSubject() {
	db := Database.Default
	data := model.Subject{
		Id:    "aaa-001",
		Title: "aaa-001",
	}
	row := db.FirstOrCreate(&data)
	if row.Error != nil {
		log.Println(row.Error)
	} else {
		log.Println(data)
	}
}
