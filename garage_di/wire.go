//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package garage_di

import (
	"github.com/google/wire"
	"github.com/gsxhnd/garage/garage_server/dao"
	"github.com/gsxhnd/garage/garage_server/handler"
	"github.com/gsxhnd/garage/garage_server/middleware"
	"github.com/gsxhnd/garage/garage_server/routes"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/utils"
)

func InitApp(path string) (*Application, error) {
	wire.Build(
		utils.UtilsSet,
		NewApplication,
		routes.NewServer,
		middleware.NewMiddleware,
		handler.HandlerSet,
		service.ServiceSet,
		dao.DaoSet,
	)
	return &Application{}, nil
}
