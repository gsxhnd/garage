package model

import (
	"database/sql"
)

type Anime struct {
	Id          uint          `json:"id"`
	Title       string        `json:"title"`
	Cover       string        `json:"cover"`
	PublishDate *sql.NullTime `json:"publish_date"`
	CreatedAt   *sql.NullTime `json:"created_at"`
	UpdatedAt   *sql.NullTime `json:"updated_at"`
}
