package handler

import "github.com/google/wire"

var HandlerSet = wire.NewSet(
	NewPingHandle,
)
