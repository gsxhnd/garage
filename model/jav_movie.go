package model

type JavMovie struct {
	Code        string `json:"code" gorm:"type:varchar(100);primaryKey"` // 番号
	Title       string `json:"title" gorm:"type:text"`                   // 名称
	PublishDate string `json:"publish_date" gorm:"type:varchar(100)"`    // 发布时间
	Director    string `json:"director" gorm:"type:varchar(100)"`        // 导演
}

func (jm *JavMovie) TableName() string {
	return "jav_movie"
}
