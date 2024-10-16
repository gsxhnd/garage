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

type AnimeHandler interface {
	GetAnime(ctx *fiber.Ctx) error
	GetAnimes(ctx *fiber.Ctx) error
	CreateAnime(ctx *fiber.Ctx) error
	DeleteAnime(ctx *fiber.Ctx) error
	UpdateAnime(ctx *fiber.Ctx) error
}

type animeHandle struct {
	valid  *validator.Validate
	svc    service.AnimeService
	logger utils.Logger
}

func NewAnimeHandler(svc service.AnimeService, v *validator.Validate, l utils.Logger) AnimeHandler {
	return &animeHandle{
		valid:  v,
		svc:    svc,
		logger: l,
	}
}

// @Summary      Create animes
// @Description  Create animes
// @Tags         anime
// @Produce      json
// @Success      200
// @Router       /anime [post]
func (h *animeHandle) CreateAnime(ctx *fiber.Ctx) error {
	var body = make([]model.Anime, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.CreateAnimes(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      Delete animes
// @Description  Delete animes
// @Tags         anime
// @Produce      json
// @Success      200
// @Router       /anime [delete]
func (h *animeHandle) DeleteAnime(ctx *fiber.Ctx) error {
	var body = make([]uint, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.DeleteAnimes(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      List animes
// @Description  Get animes
// @Tags         anime
// @Produce      json
// @Success      200
// @Router       /anime [get]
func (h *animeHandle) GetAnimes(ctx *fiber.Ctx) error {
	var p = database.Pagination{
		Limit:  uint(ctx.QueryInt("page_size", 50)),
		Offset: uint(ctx.QueryInt("page_size", 50) * ctx.QueryInt("page", 0)),
	}

	data, err := h.svc.GetAnimes(&p)

	return ctx.JSON(errno.DecodeError(err).WithData(data))
}

func (h *animeHandle) GetAnime(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}

func (h *animeHandle) UpdateAnime(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}
