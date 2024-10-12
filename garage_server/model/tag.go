package model

import (
	"time"
)

// CREATE TABLE IF NOT EXISTS "tag"
// (
//
//	"id"         INTEGER  NOT NULL UNIQUE,
//	"name"       TEXT     NOT NULL,
//	"pid"        INTEGER,
//	"created_at" DATETIME NOT NULL default CURRENT_TIMESTAMP,
//	"updated_at" DATETIME,
//	PRIMARY KEY ("id")
//
// );
type Tag struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Pid       string    `json:"p"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
