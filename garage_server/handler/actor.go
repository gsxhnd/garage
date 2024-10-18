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

type ActorHandler interface {
	CreateActors(ctx *fiber.Ctx) error
	DeleteActors(ctx *fiber.Ctx) error
	UpdateActor(ctx *fiber.Ctx) error
	GetActor(ctx *fiber.Ctx) error
	GetActors(ctx *fiber.Ctx) error
	SearchActorByName(ctx *fiber.Ctx) error
}

type actorHandle struct {
	valid  *validator.Validate
	svc    service.ActorService
	logger utils.Logger
}

func NewActorHandler(svc service.ActorService, v *validator.Validate, l utils.Logger) ActorHandler {
	return &actorHandle{
		valid:  v,
		svc:    svc,
		logger: l,
	}
}

// CreateActors implements ActorHandler.
// @Summary      Create actors
// @Description  Create actor
// @Tags         actor
// @Produce      json
// @Success      200
// @Router       /actor [post]
func (h *actorHandle) CreateActors(ctx *fiber.Ctx) error {
	var body = make([]model.Actor, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.CreateActors(body)
	return ctx.JSON(errno.DecodeError(err))
}

// DeleteActors implements ActorHandler.
// @Summary      Delete actors
// @Description  Delete actor
// @Tags         actor
// @Produce      json
// @Success      200
// @Router       /actor [delete]
func (h *actorHandle) DeleteActors(ctx *fiber.Ctx) error {
	var body = make([]uint, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.DeleteActors(body)
	return ctx.JSON(errno.DecodeError(err))
}

// UpdateActor implements ActorHandler.
// @Summary      Update a actor by id
// @Description  Update a actor by id
// @Tags         actor
// @Accept       json
// @Produce      json
// @Param        tag body model.Actor true "Actor object"
// @Success      200 {object} errno.errno
// @Router       /actor [put]
func (h *actorHandle) UpdateActor(ctx *fiber.Ctx) error {
	var body model.Actor
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Struct(body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}
	h.svc.UpdateActor(&body)

	return ctx.JSON(errno.OK)
}

// GetActors implements ActorHandler.
// @Summary      Get actors
// @Description  Get actors List
// @Tags         actor
// @Produce      json
// @Success      200  {object}   errno.errno{data=[]model.Actor}
// @Router       /actor [get]
func (h *actorHandle) GetActors(ctx *fiber.Ctx) error {
	var p = database.Pagination{
		Limit:  uint(ctx.QueryInt("page_size", 50)),
		Offset: uint(ctx.QueryInt("page_size", 50) * ctx.QueryInt("page", 0)),
	}

	data, err := h.svc.GetActors(&p)

	return ctx.JSON(errno.DecodeError(err).WithData(data))
}

// GetActor implements ActorHandler.
// TODO: get actor movies
func (h *actorHandle) GetActor(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// @Summary      Search actors
// @Description  Search actors List
// @Tags         actor
// @Param        name    query     string  false  "name search by name"
// @Produce      json
// @Success      200  {object}   errno.errno{data=[]model.Actor}
// @Router       /actor/search [get]
func (h *actorHandle) SearchActorByName(ctx *fiber.Ctx) error {
	data, err := h.svc.SearchActorByName(ctx.Query("name"))
	if err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}
	return ctx.JSON(errno.OK.WithData(data))
}
