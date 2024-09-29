package handler

import (
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
	// validator *validator.Validate
	// svc service.PingService
}

func NewMovieHandler(svc service.PingService) MovieHandler {
	return &movieHandle{
		// svc: svc,
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
