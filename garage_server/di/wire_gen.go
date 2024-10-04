// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/gsxhnd/garage/garage_server/db"
	"github.com/gsxhnd/garage/garage_server/handler"
	"github.com/gsxhnd/garage/garage_server/middleware"
	"github.com/gsxhnd/garage/garage_server/router"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/garage_server/storage"
	"github.com/gsxhnd/garage/utils"
)

// Injectors from wire.go:

func InitApp() (*Application, error) {
	config, err := utils.NewConfig()
	if err != nil {
		return nil, err
	}
	logger := utils.NewLogger(config)
	middlewareMiddleware := middleware.NewMiddleware(logger)
	database, err := db.NewDatabase(config, logger)
	if err != nil {
		return nil, err
	}
	storageStorage, err := storage.NewStorage(config)
	if err != nil {
		return nil, err
	}
	pingService := service.NewPingService(logger, database, storageStorage)
	pingHandler := handler.NewPingHandler(pingService)
	movieService := service.NewMovieService(logger, database)
	validate := utils.NewValidator()
	movieHandler := handler.NewMovieHandler(movieService, validate)
	starService := service.NewStarService(logger, database)
	starHandler := handler.NewStarHandler(starService, validate)
	handlerHandler := handler.Handler{
		PingHandler:  pingHandler,
		MovieHandler: movieHandler,
		StarHandler:  starHandler,
	}
	routerRouter, err := router.NewRouter(config, middlewareMiddleware, handlerHandler)
	if err != nil {
		return nil, err
	}
	application := NewApplication(routerRouter)
	return application, nil
}
