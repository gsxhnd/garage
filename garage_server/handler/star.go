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

type StarHandler interface {
	GetStar(ctx *fiber.Ctx) error
	GetStars(ctx *fiber.Ctx) error
	CreateStars(ctx *fiber.Ctx) error
	DeleteStars(ctx *fiber.Ctx) error
	UpdateStar(ctx *fiber.Ctx) error
}

type starHandle struct {
	valid  *validator.Validate
	svc    service.StarService
	logger utils.Logger
}

func NewStarHandler(svc service.StarService, v *validator.Validate, l utils.Logger) StarHandler {
	return &starHandle{
		valid:  v,
		svc:    svc,
		logger: l,
	}
}

// CreateStars implements StarHandler.
func (h *starHandle) CreateStars(ctx *fiber.Ctx) error {
	var body = make([]model.Star, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.CreateStars(body)
	return ctx.JSON(errno.DecodeError(err))
}

// DeleteStars implements StarHandler.
func (h *starHandle) DeleteStars(ctx *fiber.Ctx) error {
	var body = make([]uint, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.DeleteStars(body)
	return ctx.JSON(errno.DecodeError(err))
}

// GetStar implements StarHandler.
func (h *starHandle) GetStar(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetStars implements StarHandler.
func (h *starHandle) GetStars(ctx *fiber.Ctx) error {
	var p = database.Pagination{
		Limit:  uint(ctx.QueryInt("page_size", 50)),
		Offset: uint(ctx.QueryInt("page_size", 50) * ctx.QueryInt("page", 0)),
	}

	data, err := h.svc.GetStars(&p)

	return ctx.JSON(errno.DecodeError(err).WithData(data))
}

// UpdateStar implements StarHandler.
func (h *starHandle) UpdateStar(ctx *fiber.Ctx) error {
	panic("unimplemented")
}
