package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gsxhnd/garage/garage_server/handler"
	"github.com/gsxhnd/garage/garage_server/middleware"
)

type Routes struct {
	Engine      *gin.Engine
	Middleware  middleware.Middlewarer
	PingHandler handler.PingHandler
	// DemoHandle handler.DemoHandler
}

func (r *Routes) Init() {
	// r.Engine.Use(r.Middleware.RequestLog())

	rootRoutes := r.Engine.Group("/")
	rootRoutes.GET("/ping", r.PingHandler.Ping)
	// r.newDemoRoute()
}

var RouteSet = wire.NewSet(
	gin.New,
	middleware.NewMiddleware,
	handler.NewPingHandle,
	// handler.NewDemoHandle,

	wire.Struct(new(Routes),
		"Engine",
		"Middleware",
		"PingHandler",
		// "DemoHandle",
	),
)
