package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/errno"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/utils"
)

type MovieActorHandler interface {
	GetMovieActors(ctx *fiber.Ctx) error
	CreateMovieActors(ctx *fiber.Ctx) error
	DeleteMovieActors(ctx *fiber.Ctx) error
	UpdateMovieActor(ctx *fiber.Ctx) error
}

type movieActorHandle struct {
	valid  *validator.Validate
	svc    service.MovieActorService
	logger utils.Logger
}

func NewMovieActorHandler(svc service.MovieActorService, v *validator.Validate, l utils.Logger) MovieActorHandler {
	return &movieActorHandle{
		valid:  v,
		svc:    svc,
		logger: l,
	}
}

// @Summary      Get movie actor
// @Description  Get movie actor
// @Tags         movie_actor
// @Produce      json
// @Success      200
// @Router       /movie_actor [post]
func (h *movieActorHandle) CreateMovieActors(ctx *fiber.Ctx) error {
	var body = make([]model.MovieActor, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.CreateMovieActors(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      Delete movie actor
// @Description  Delete movie actor
// @Tags         movie_actor
// @Produce      json
// @Success      200
// @Router       /movie_actor [delete]
func (h *movieActorHandle) DeleteMovieActors(ctx *fiber.Ctx) error {
	var body = make([]uint, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.DeleteMovieActors(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      Get movie actor
// @Description  Get movie actor
// @Tags         movie_actor
// @Produce      json
// @Success      200
// @Router       /movie_actor [get]
func (h *movieActorHandle) GetMovieActors(ctx *fiber.Ctx) error {
	data, err := h.svc.GetMovieActors()

	return ctx.JSON(errno.DecodeError(err).WithData(data))
}

func (h *movieActorHandle) UpdateMovieActor(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}
