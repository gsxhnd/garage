package model

import "time"

type Tag struct {
	Id        uint       `json:"id" validate:"required"`
	Name      string     `json:"name" validate:"required"`
	Pid       uint       `json:"pid"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
