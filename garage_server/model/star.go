package model

import "time"

type Star struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	AliasName string    `json:"alias_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
