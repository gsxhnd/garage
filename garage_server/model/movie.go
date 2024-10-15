package model

import (
	"database/sql"
)

type Movie struct {
	Id             uint            `json:"id"`
	Code           string          `json:"code" validate:"required"`
	Title          string          `json:"title" validate:"required"`
	Cover          *sql.NullTime   `json:"cover"`
	PublishDate    *sql.NullTime   `json:"publish_date"`
	Director       *sql.NullString `json:"director"`
	ProduceCompany *sql.NullString `json:"produce_company"`
	PublishCompany *sql.NullString `json:"publish_company"`
	Series         *sql.NullString `json:"series"`
	CreatedAt      *sql.NullTime   `json:"created_at"`
	UpdatedAt      *sql.NullTime   `json:"updated_at"`
}
