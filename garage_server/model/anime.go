package model

import (
	"database/sql"
	"time"
)

type Anime struct {
	Id          uint          `json:"id"`
	Title       string        `json:"title"`
	Cover       string        `json:"cover"`
	PublishDate *sql.NullTime `json:"publish_date"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
