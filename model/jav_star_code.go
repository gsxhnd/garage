package model

type JavStarCode struct {
	JavId          uint64 `json:"jav_id" gorm:"primaryKey;autoIncrement;"`
	JavbusCode     string `json:"javbus_code" gorm:"type:varchar(100)"`
	JavlibraryCode string `json:"javlibrary_code" gorm:"type:varchar(100)"`
}

func (js *JavStarCode) TableName() string {
	return "jav_star_code"
}
