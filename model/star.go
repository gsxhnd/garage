package model

type Star struct {
	Id    uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Image string `json:"image"`
	TimestampModel
}
