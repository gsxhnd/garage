package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/garage_server/errno"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/utils"
)

type MovieHandler interface {
	GetMovie(ctx *fiber.Ctx) error
	GetMovies(ctx *fiber.Ctx) error
	CreateMovies(ctx *fiber.Ctx) error
	DeleteMovies(ctx *fiber.Ctx) error
	UpdateMovie(ctx *fiber.Ctx) error
}

type movieHandle struct {
	valid  *validator.Validate
	svc    service.MovieService
	logger utils.Logger
}

func NewMovieHandler(svc service.MovieService, v *validator.Validate, l utils.Logger) MovieHandler {
	return &movieHandle{
		valid:  v,
		svc:    svc,
		logger: l,
	}
}

func (h *movieHandle) CreateMovies(ctx *fiber.Ctx) error {
	var body = make([]model.Movie, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.CreateMovies(body)
	return ctx.JSON(errno.DecodeError(err))
}

func (h *movieHandle) DeleteMovies(ctx *fiber.Ctx) error {
	var body = make([]uint, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.DeleteMovies(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Description  Get movies
// @Produce      json
// @Success      200
// @Router       /ping [get]
func (h *movieHandle) GetMovies(ctx *fiber.Ctx) error {
	var p = database.Pagination{
		Limit:  uint(ctx.QueryInt("page_size", 50)),
		Offset: uint(ctx.QueryInt("page_size", 50) * ctx.QueryInt("page", 0)),
	}

	data, err := h.svc.GetMovies(&p)

	return ctx.JSON(errno.DecodeError(err).WithData(data))
}

func (h *movieHandle) GetMovie(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}

func (h *movieHandle) UpdateMovie(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}
