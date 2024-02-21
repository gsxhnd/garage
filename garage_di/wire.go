//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package garage_di

import (
	"github.com/google/wire"
	"github.com/gsxhnd/garage/garage_server/routes"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/utils"
)

func InitApp() (*Application, error) {
	wire.Build(
		NewApplication,
		routes.RouteSet,
		utils.UtilsSet,
		service.ServiceSet,
		// dao.DaoSet,
	)
	return &Application{}, nil
}
