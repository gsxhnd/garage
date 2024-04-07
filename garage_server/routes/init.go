package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
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
	ffmpegGroup.GET("/videos") //Get all video list
	ffmpegGroup.POST("/task")  //Get all video list

	taskGroup := g.Group("/task")
	taskGroup.GET("/")
	taskGroup.POST("/") // Create Task

	return &Routers{
		Engine: g,
	}
}

func NewServer(cfg *utils.Config, h handler.Handler) *fiber.App {
	app := fiber.New()
	rootGroup := app.Group("/")
	rootGroup.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	return app
}
