package model

import "time"

type TimestampModel struct {
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
