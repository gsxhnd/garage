package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gsxhnd/garage/garage_server/handler"
	"github.com/gsxhnd/garage/garage_server/middleware"
	"github.com/gsxhnd/garage/garage_web"
	"github.com/gsxhnd/garage/utils"
)

type Router interface {
	Run() error
}

type router struct {
	cfg *utils.Config
	app *fiber.App
	h   handler.Handler
	m   middleware.Middleware
}

// @title           Garage API
// @version         1
// @description     This is a sample server celler server.
// @license.name  MIT
// @license.url   https://opensource.org/license/mit
// @host      localhost:8080
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
func NewRouter(cfg *utils.Config, m middleware.Middleware, h handler.Handler) (Router, error) {
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

	api := r.app.Group("/api/v1")
	// api.Get("/crawl")
	api.Post("/jav/movie", r.h.MovieHandler.CreateMovies)
	api.Delete("/jav/movie", r.h.MovieHandler.DeleteMovies)
	api.Put("/jav/movie/:code", r.h.MovieHandler.UpdateMovie)
	api.Get("/jav/movie", r.h.MovieHandler.GetMovies)
	api.Post("/jav/star", r.h.StarHandler.CreateStars)
	api.Delete("/jav/star", r.h.StarHandler.DeleteStars)
	api.Put("/jav/star/:code", r.h.StarHandler.UpdateStar)
	api.Get("/jav/star", r.h.StarHandler.GetStars)

	r.app.Use("/*", filesystem.New(filesystem.Config{
		Root:       http.FS(garage_web.Content),
		PathPrefix: "dist",
		Browse:     true,
	}))

	r.app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	return r.app.Listen(r.cfg.Listen)
}