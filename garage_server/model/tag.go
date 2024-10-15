package model

import (
	"database/sql"
)

type Tag struct {
	Id        uint          `json:"id"`
	Name      string        `json:"name"`
	Pid       uint          `json:"pid"`
	CreatedAt *sql.NullTime `json:"created_at"`
	UpdatedAt *sql.NullTime `json:"updated_at"`
}
