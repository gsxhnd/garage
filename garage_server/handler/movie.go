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
	CreateMovies(ctx *fiber.Ctx) error
	DeleteMovies(ctx *fiber.Ctx) error
	UpdateMovie(ctx *fiber.Ctx) error
	GetMovies(ctx *fiber.Ctx) error
	GetMovieInfo(ctx *fiber.Ctx) error
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

// @Summary      Create movies
// @Description  Create movies
// @Tags         movie
// @Produce      json
// @Success      200  {object}  errno.errno
// @Router       /movie [post]
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

// @Summary      Delete movies
// @Description  Delete movies
// @Tags         movie
// @Accept       json
// @Produce      json
// @Param        default body []uint true "default"
// @Success      200 {object} errno.errno{data=nil}
// @Router       /movie [delete]
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

// @Summary      Get movies
// @Description  Get movies
// @Tags         movie
// @Produce      json
// @Param        page_size query int false "int valid" default(50)
// @Param        page query int false "int valid" default(1)
// @Success      200 {object} errno.errno{data=nil}
// @Router       /movie [get]
func (h *movieHandle) GetMovies(ctx *fiber.Ctx) error {
	var p = database.Pagination{
		Limit:  uint(ctx.QueryInt("page_size", 50)),
		Offset: uint(ctx.QueryInt("page_size", 50) * (ctx.QueryInt("page", 1) - 1)),
	}

	if err := h.valid.Struct(p); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	data, err := h.svc.GetMovies(&p)

	return ctx.JSON(errno.DecodeError(err).WithData(data))
}

// @Summary      Get movies
// @Description  Get movies
// @Tags         movie
// @Produce      json
// @Param        code path string true "movie code"
// @Success      200 {object} errno.errno{data=model.MovieInfo}
// @Router       /movie/:code [get]
func (h *movieHandle) GetMovieInfo(ctx *fiber.Ctx) error {
	code := ctx.Params("code", "")
	data, err := h.svc.GetMovieInfo(code)
	return ctx.JSON(errno.DecodeError(err).WithData(data))
}

func (h *movieHandle) UpdateMovie(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}
