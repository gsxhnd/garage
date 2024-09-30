package utils

import "github.com/google/wire"

var UtilsSet = wire.NewSet(
	NewLogger, NewConfig, NewValidator,
)
