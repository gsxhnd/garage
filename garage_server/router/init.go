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
	cfg    *utils.Config
	app    *fiber.App
	logger utils.Logger
	h      handler.Handler
	m      middleware.Middleware
}

// @title           Garage API
// @version         1
// @description     This is a sample server celler server.
// @license.name  MIT
// @license.url   https://opensource.org/license/mit
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
func NewRouter(cfg *utils.Config, l utils.Logger, m middleware.Middleware, h handler.Handler) (Router, error) {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes:     cfg.Mode == "dev",
		DisableStartupMessage: cfg.Mode == "prod",
		Prefork:               false,
	})

	return &router{
		cfg:    cfg,
		app:    app,
		logger: l,
		h:      h,
		m:      m,
	}, nil
}

func (r *router) Run() error {
	// r.app.Use(r.m.RequestLog)
	r.app.Get("/ping", r.h.PingHandler.Ping)

	api := r.app.Group("/api/v1")
	// movie api
	api.Post("/movie", r.h.MovieHandler.CreateMovies)
	api.Delete("/movie", r.h.MovieHandler.DeleteMovies)
	api.Put("/movie/:code", r.h.MovieHandler.UpdateMovie)
	api.Get("/movie", r.h.MovieHandler.GetMovies)
	api.Get("/movie/:code", r.h.MovieHandler.GetMovieInfo)
	// actor api
	api.Post("/actor", r.h.ActorHandler.CreateActors)
	api.Delete("/actor", r.h.ActorHandler.DeleteActors)
	api.Put("/actor", r.h.ActorHandler.UpdateActor)
	api.Get("/actor", r.h.ActorHandler.GetActors)
	api.Get("/actor/search", r.h.ActorHandler.SearchActorByName)
	// tag
	api.Post("/tag", r.h.TagHandler.CreateTag)
	api.Delete("/tag", r.h.TagHandler.DeleteTag)
	api.Put("/tag", r.h.TagHandler.UpdateTag)
	api.Get("/tag", r.h.TagHandler.GetTags)
	// movie actor
	api.Post("/movie_actor", r.h.MovieActorHandle.CreateMovieActors)
	api.Delete("/movie_actor", r.h.MovieActorHandle.DeleteMovieActors)
	api.Get("/movie_actor", r.h.MovieActorHandle.GetMovieActors)
	// movie tag
	api.Post("/movie_tag", r.h.MovieActorHandle.CreateMovieActors)
	api.Delete("/movie_tag", r.h.MovieActorHandle.DeleteMovieActors)
	api.Get("/movie_tag", r.h.MovieActorHandle.GetMovieActors)
	// anime
	api.Post("/anime", r.h.AnimeHandler.CreateAnime)
	api.Delete("/anime", r.h.AnimeHandler.DeleteAnime)
	api.Put("/anime", r.h.AnimeHandler.UpdateAnime)
	api.Get("/anime", r.h.AnimeHandler.GetAnimes)

	img := r.app.Group("/api/v1/img")
	img.Get("/movie/:id", r.h.ImageHandler.GetMovieImage)
	img.Get("/actor/:id", r.h.ImageHandler.GetActorImage)

	r.app.Use("/*", filesystem.New(filesystem.Config{
		Root:       http.FS(garage_web.Content),
		PathPrefix: "dist",
		Browse:     true,
	}))

	r.app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	r.logger.Infof("Server actort listening")

	return r.app.Listen(r.cfg.Listen)
}
