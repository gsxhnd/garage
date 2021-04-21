package model

import "time"

type JavMovie struct {
	Code        string    `json:"code" gorm:"type:varchar(100);primaryKey"` // 番号
	Title       string    `json:"title" gorm:"type:text"`                   // 名称
	PublishDate time.Time `json:"publish_date" gorm:"type:datetime"`        // 发布时间
	Director    string    `json:"director" gorm:"type:varchar(100)"`        // 导演
}

func (jm *JavMovie) TableName() string {
	return "jav_movie"
}
