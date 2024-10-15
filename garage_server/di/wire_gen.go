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
	driver, err := db.NewDatabase(config, logger)
	if err != nil {
		return nil, err
	}
	storageStorage, err := storage.NewStorage(config)
	if err != nil {
		return nil, err
	}
	pingService := service.NewPingService(logger, driver, storageStorage)
	pingHandler := handler.NewPingHandler(pingService)
	movieService := service.NewMovieService(logger, driver)
	validate := utils.NewValidator()
	movieHandler := handler.NewMovieHandler(movieService, validate, logger)
	movieStarService := service.NewMovieStarService(logger, driver)
	movieStarHandler := handler.NewMovieStarHandler(movieStarService, validate, logger)
	movieTagService := service.NewMovieTagService(logger, driver)
	movieTagHandler := handler.NewMovieTagHandler(movieTagService, validate, logger)
	starService := service.NewStarService(logger, driver)
	starHandler := handler.NewStarHandler(starService, validate, logger)
	imageHandler := handler.NewImageHandler(validate, storageStorage, logger)
	tagService := service.NewTagService(logger, driver)
	tagHandler := handler.NewTagHandler(tagService, validate, logger)
	animeService := service.NewAnimeService(logger, driver)
	animeHandler := handler.NewAnimeHandler(animeService, validate)
	handlerHandler := handler.Handler{
		PingHandler:     pingHandler,
		MovieHandler:    movieHandler,
		MovieStarHandle: movieStarHandler,
		MovieTagHandler: movieTagHandler,
		StarHandler:     starHandler,
		ImageHandler:    imageHandler,
		TagHandler:      tagHandler,
		AnimeHandler:    animeHandler,
	}
	routerRouter, err := router.NewRouter(config, logger, middlewareMiddleware, handlerHandler)
	if err != nil {
		return nil, err
	}
	application := NewApplication(routerRouter)
	return application, nil
}
