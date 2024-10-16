package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/errno"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/utils"
)

type MovieTagHandler interface {
	GetMovieTag(ctx *fiber.Ctx) error
	GetMovieTags(ctx *fiber.Ctx) error
	CreateMovieTags(ctx *fiber.Ctx) error
	DeleteMovieTags(ctx *fiber.Ctx) error
	UpdateMovieTag(ctx *fiber.Ctx) error
}

type movieTagHandle struct {
	valid  *validator.Validate
	svc    service.MovieTagService
	logger utils.Logger
}

func NewMovieTagHandler(svc service.MovieTagService, v *validator.Validate, l utils.Logger) MovieTagHandler {
	return &movieTagHandle{
		valid:  v,
		svc:    svc,
		logger: l,
	}
}

// @Summary      Create movie tags
// @Description  Create movie tags
// @Tags         movie_tag
// @Produce      json
// @Success      200  {object}   errno.errno
// @Router       /movie_tag [post]
func (h *movieTagHandle) CreateMovieTags(ctx *fiber.Ctx) error {
	var body = make([]model.MovieTag, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.CreateMovieTags(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      Delete movie tags
// @Description  Delete movie tags
// @Tags         movie_tag
// @Produce      json
// @Success      200  {object}   errno.errno
// @Router       /movie_tag [delete]
func (h *movieTagHandle) DeleteMovieTags(ctx *fiber.Ctx) error {
	var body = make([]uint, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.DeleteMovieTags(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      Get movie tags
// @Description  Get movie tags
// @Tags         movie_tag
// @Produce      json
// @Success      200  {object}   errno.errno{data=[]model.MovieTag}
// @Router       /movie_tag [get]
func (h *movieTagHandle) GetMovieTags(ctx *fiber.Ctx) error {
	data, err := h.svc.GetMovieTags()

	return ctx.JSON(errno.DecodeError(err).WithData(data))
}

func (h *movieTagHandle) GetMovieTag(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}

func (h *movieTagHandle) UpdateMovieTag(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}
