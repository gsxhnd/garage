package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/errno"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/utils"
)

type MovieStarHandler interface {
	GetMovieStars(ctx *fiber.Ctx) error
	CreateMovieStars(ctx *fiber.Ctx) error
	DeleteMovieStars(ctx *fiber.Ctx) error
	UpdateMovieStar(ctx *fiber.Ctx) error
}

type movieStarHandle struct {
	valid  *validator.Validate
	svc    service.MovieStarService
	logger utils.Logger
}

func NewMovieStarHandler(svc service.MovieStarService, v *validator.Validate, l utils.Logger) MovieStarHandler {
	return &movieStarHandle{
		valid:  v,
		svc:    svc,
		logger: l,
	}
}

// @Summary      Get movie star
// @Description  Get movie star
// @Tags         movie_star
// @Produce      json
// @Success      200
// @Router       /movie_star [post]
func (h *movieStarHandle) CreateMovieStars(ctx *fiber.Ctx) error {
	var body = make([]model.MovieStar, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.CreateMovieStars(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      Delete movie star
// @Description  Delete movie star
// @Tags         movie_star
// @Produce      json
// @Success      200
// @Router       /movie_star [delete]
func (h *movieStarHandle) DeleteMovieStars(ctx *fiber.Ctx) error {
	var body = make([]uint, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.DeleteMovieStars(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      Get movie star
// @Description  Get movie star
// @Tags         movie_star
// @Produce      json
// @Success      200
// @Router       /movie_star [get]
func (h *movieStarHandle) GetMovieStars(ctx *fiber.Ctx) error {
	data, err := h.svc.GetMovieStars()

	return ctx.JSON(errno.DecodeError(err).WithData(data))
}

func (h *movieStarHandle) UpdateMovieStar(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}
