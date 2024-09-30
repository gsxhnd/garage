//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"github.com/google/wire"
	"github.com/gsxhnd/garage/garage_server/db"
	"github.com/gsxhnd/garage/garage_server/handler"
	"github.com/gsxhnd/garage/garage_server/middleware"
	"github.com/gsxhnd/garage/garage_server/router"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/garage_server/storage"
	"github.com/gsxhnd/garage/utils"
)

func InitApp() (*Application, error) {
	wire.Build(
		utils.UtilsSet,
		NewApplication,
		router.NewRouter,
		middleware.NewMiddleware,
		handler.HandlerSet,
		service.ServiceSet,
		storage.StorageSet,
		db.DBSet,
	)
	return &Application{}, nil
}
