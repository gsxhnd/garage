package db

import "github.com/gsxhnd/garage/utils"

type TaskDao interface{}

type taskDao struct {
	logger utils.Logger
	db     *Database
}

func NewTaskDao(db *Database, l utils.Logger) TaskDao {
	return &taskDao{
		logger: l,
		db:     db,
	}
}
