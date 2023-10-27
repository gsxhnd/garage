//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"github.com/google/wire"
)

func InitLog() {
	wire.Build()
}
