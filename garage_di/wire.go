//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package garage_di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gsxhnd/garage/garage_server/handler"
	"github.com/gsxhnd/garage/garage_server/middleware"
	"github.com/gsxhnd/garage/garage_server/routes"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/utils"
)

func InitApp() (*Application, error) {
	wire.Build(
		gin.New,
		NewApplication,
		routes.NewRouter,
		middleware.NewMiddleware,
		utils.UtilsSet,
		handler.HandlerSet,
		service.ServiceSet,
		// dao.DaoSet,
	)
	return &Application{}, nil
}
