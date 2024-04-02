package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/garage/garage_server/handler"
	"github.com/gsxhnd/garage/garage_server/middleware"
	"github.com/gsxhnd/garage/utils"
)

type Routers struct {
	Engine *gin.Engine
}

func NewRouter(cfg *utils.Config, g *gin.Engine, h handler.Handler, m middleware.Middlewarer) *Routers {
	g.Use(m.RequestLog())
	rootGroup := g.Group("/")
	rootGroup.GET("/ping", h.PingHandler.Ping)
	rootGroup.GET("/ws", h.WebsocketHandler.Ws)

	ffmpegGroup := g.Group("/ffmpeg")
	ffmpegGroup.GET("/videos")  //Get all video list
	ffmpegGroup.POST("/task") //Get all video list

	taskGroup := g.Group("/task")
	taskGroup.GET("/")
	taskGroup.POST("/") // Create Task

	return &Routers{
		Engine: g,
	}
}
