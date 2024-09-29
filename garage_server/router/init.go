package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/handler"
	"github.com/gsxhnd/garage/garage_server/middleware"
	"github.com/gsxhnd/garage/utils"
)

type Router interface {
	Run() error
}

type router struct {
	cfg *utils.Config
	app *fiber.App
	h   handler.Handler
	m   middleware.Middlewarer
}

// @title           Tenhou API
// @version         1
// @description     This is a sample server celler server.
// @license.name  MIT
// @license.url   https://opensource.org/license/mit
// @host      localhost:8080
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
func NewRouter(cfg *utils.Config, m middleware.Middlewarer, h handler.Handler) (Router, error) {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes:     cfg.Mode == "dev",
		DisableStartupMessage: cfg.Mode == "prod",
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
	// r.app.Use(r.m.RequestLog)
	r.app.Get("/ping", r.h.PingHandler.Ping)

	// api := r.app.Group("/api/v1")
	// api.Get("/log/:log_id", r.h.LogHandler.GetLogInfoByLogId)
	// api.Get("/paifu")
	// api.Get("/paifu/:log_id")

	r.app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	return r.app.Listen(r.cfg.Listen)
}
