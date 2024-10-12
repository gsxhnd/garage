package model

import (
	"database/sql"
	"time"
)

type Movie struct {
	Id             uint          `json:"id"`
	Code           string        `json:"code"`
	Title          string        `json:"title"`
	PublishDate    *sql.NullTime `json:"publish_date"`
	Director       string        `json:"director"`
	ProduceCompany string        `json:"produce_company"`
	PublishCompany string        `json:"publish_company"`
	Series         string        `json:"series"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}
