package routes

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/handler"
	"github.com/gsxhnd/garage/garage_server/middleware"
	"github.com/gsxhnd/garage/utils"
)

type Router interface {
	Run() error
	EnableWeb() error
}

type router struct {
	cfg *utils.Config
	app *fiber.App
	h   handler.Handler
	m   middleware.Middlewarer
}

func NewRouter(cfg *utils.Config, h handler.Handler, m middleware.Middlewarer) (Router, error) {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes:     cfg.Dev,
		DisableStartupMessage: !cfg.Dev,
		Prefork:               false,
	})

	return &router{
		cfg: cfg,
		app: app,
		h:   h,
		m:   m,
	}, nil
}

func (r *router) Run() error {
	r.app.Use("/ws", r.m.Websocket)
	ws := r.app.Group("/ws")
	ws.Get("", websocket.New(r.h.WebsocketHandler.Ws))

	r.ApiInit()

	if err := r.EnableWeb(); err != nil {
		return err
	}

	r.app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return r.app.Listen(r.cfg.Port)
}
