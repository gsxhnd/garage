package routes

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	garage_ui "github.com/gsxhnd/garage/garage-ui"
	"github.com/gsxhnd/garage/garage_server/handler"
	"github.com/gsxhnd/garage/garage_server/middleware"
	"github.com/gsxhnd/garage/utils"
)

type Routers struct {
	app *fiber.App
}

func NewServer(cfg *utils.Config, h handler.Handler, m middleware.Middlewarer) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes:     true,
		DisableStartupMessage: false,
	})

	dist, err := fs.Sub(garage_ui.Web, "dist")
	if err != nil {
		log.Fatalf("dist file server")
		return nil, err
	}

	app.Use("/ws", m.Websocket)
	ws := app.Group("/ws")
	ws.Get("", websocket.New(h.WebsocketHandler.Ws))

	api := app.Group("/api")
	api.Get("/ping", h.RootHandler.Ping)

	ffmpegGroup := api.Group("/ffmpeg")
	ffmpegGroup.Get("/videos", h.RootHandler.Ping)

	webG := app.Group("/")
	webG.All("*", filesystem.New(filesystem.Config{
		Root:  http.FS(dist),
		Index: "index.html",
	}))

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return app, nil
}
