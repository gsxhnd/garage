package service

import "github.com/google/wire"

var ServiceSet = wire.NewSet(
	NewTaskService,
)
