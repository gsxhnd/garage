package model

import (
	"time"
)

type Actor struct {
	Id        uint       `json:"id" validate:"required"`
	Name      string     `json:"name" validate:"required"`
	AliasName *string    `json:"alias_name"`
	Cover     *string    `json:"cover"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
