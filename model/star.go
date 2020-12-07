package model

type Star struct {
	Id    uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"type:varchar(100)"`
	Image string `json:"image" gorm:"type:varchar(100)"`
	TimestampModel
}
