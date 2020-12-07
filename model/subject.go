package model

import "time"

type Subject struct {
	Id          string    `json:"id" gorm:"type:varchar(100);primaryKey"`
	Title       string    `json:"title"`
	PublishTime time.Time `json:"publish_time"`
	TimestampModel
}

func (s Subject) TableName() string {
	return "subject"
}
