package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/service"
)

type MovieHandler interface {
	GetMovie(ctx *fiber.Ctx) error
	GetMovies(ctx *fiber.Ctx) error
	CreateMovies(ctx *fiber.Ctx) error
	DeleteMovies(ctx *fiber.Ctx) error
	UpdateMovie(ctx *fiber.Ctx) error
}

type movieHandle struct {
	valid *validator.Validate
	svc   service.MovieService
}

func NewMovieHandler(svc service.MovieService, v *validator.Validate) MovieHandler {
	return &movieHandle{
		valid: v,
		svc:   svc,
	}
}

// @Description  ping serivce working, db connect
// @Produce      json
// @Success      200
// @Router       /ping [get]
func (h *movieHandle) GetMovies(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}

func (h *movieHandle) GetMovie(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}

func (h *movieHandle) CreateMovies(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}

func (h *movieHandle) DeleteMovies(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}

func (h *movieHandle) UpdateMovie(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}
