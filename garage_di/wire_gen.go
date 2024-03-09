// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package garage_di

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/garage/garage_server/handler"
	"github.com/gsxhnd/garage/garage_server/middleware"
	"github.com/gsxhnd/garage/garage_server/routes"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/utils"
)

// Injectors from wire.go:

func InitApp() (*Application, error) {
	engine := gin.New()
	config, err := utils.NewConfig()
	if err != nil {
		return nil, err
	}
	logger := utils.NewLogger(config)
	middlewarer := middleware.NewMiddleware(logger)
	testService := service.NewTestService(logger)
	pingHandler := handler.NewPingHandle(logger, testService)
	websocketHandler := handler.NewWebsocketHandler(logger)
	routesRoutes := &routes.Routes{
		Engine:           engine,
		Middleware:       middlewarer,
		PingHandler:      pingHandler,
		WebsocketHandler: websocketHandler,
	}
	application := NewApplication(routesRoutes)
	return application, nil
}
