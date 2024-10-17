package model

import (
	"database/sql"
)

type Actor struct {
	Id        uint            `json:"id"`
	Name      string          `json:"name"`
	AliasName *sql.NullString `json:"alias_name"`
	CreatedAt *sql.NullTime   `json:"created_at"`
	UpdatedAt *sql.NullTime   `json:"updated_at"`
}
