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
	CreateMovieTags(ctx *fiber.Ctx) error
	DeleteMovieTags(ctx *fiber.Ctx) error
	GetMovieTags(ctx *fiber.Ctx) error
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
// @Param        default body []model.MovieTag true "default"
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
// @Param        default body []uint true "default"
// @Success      200 {object} errno.errno{data=nil}
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

// @Summary      Get movie id tags
// @Description  Get movie id tags
// @Tags         movie_tag
// @Produce      json
// @Param        movie_id path uint true "movie id"
// @Success      200  {object}   errno.errno{data=[]model.MovieTag}
// @Router       /movie_tag/{movie_id} [get]
func (h *movieTagHandle) GetMovieTags(ctx *fiber.Ctx) error {
	p := struct {
		MovieId uint `params:"movie_id"`
	}{}

	if err := ctx.ParamsParser(&p); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	data, err := h.svc.GetMovieTags(p.MovieId)
	return ctx.JSON(errno.DecodeError(err).WithData(data))
}
