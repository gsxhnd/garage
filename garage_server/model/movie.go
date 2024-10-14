package model

import (
	"database/sql"
	"time"
)

type Movie struct {
	Id             uint            `json:"id"`
	Code           string          `json:"code" validate:"required"`
	Title          string          `json:"title" validate:"required"`
	PublishDate    *sql.NullTime   `json:"publish_date"`
	Director       *sql.NullString `json:"director"`
	ProduceCompany *sql.NullString `json:"produce_company"`
	PublishCompany *sql.NullString `json:"publish_company"`
	Series         *sql.NullString `json:"series"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
