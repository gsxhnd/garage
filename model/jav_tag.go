package model

type JavTag struct {
	Id    uint64 `json:"id" gorm:"primaryKey;autoIncrement;"`
	Label string `json:"label" gorm:"type:varchar(100)"`
}

func (jt *JavTag) TableName() string {
	return "jav_tag"
}
