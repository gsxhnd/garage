package model

type JavStar struct {
	Id    uint64 `json:"id" gorm:"primaryKey;autoIncrement;"`
	Name  string `json:"name" gorm:"type:varchar(100)"` // 名字
	Image string `json:"image" gorm:"type:varchar(100)"`
}

func (js *JavStar) TableName() string {
	return "jav_star"
}
