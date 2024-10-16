package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/errno"
	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/gsxhnd/garage/garage_server/service"
	"github.com/gsxhnd/garage/utils"
)

type TagHandler interface {
	GetTag(ctx *fiber.Ctx) error
	GetTags(ctx *fiber.Ctx) error
	CreateTag(ctx *fiber.Ctx) error
	DeleteTag(ctx *fiber.Ctx) error
	UpdateTag(ctx *fiber.Ctx) error
}

type tagHandle struct {
	valid  *validator.Validate
	svc    service.TagService
	logger utils.Logger
}

func NewTagHandler(svc service.TagService, v *validator.Validate, l utils.Logger) TagHandler {
	return &tagHandle{
		valid:  v,
		svc:    svc,
		logger: l,
	}
}

// @Summary      Get a single tag by ID
// @Description  Get a single tag by ID
// @Tags         tag
// @Produce      json
// @Param        id   path      string  true  "Tag ID"
// @Success      200  {object}   errno.errno
// @Router       /tag/{id} [get]
func (h *tagHandle) GetTag(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(errno.InternalServerError)
	}

	// tag, err := h.svc.GetTag(id)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(errno.DecodeError(err))
	// }

	return ctx.JSON(errno.OK)
}

// @Summary      Get all tags
// @Description  Get all tags
// @Tags         tag
// @Produce      json
// @Success      200  {object}   errno.errno{data=[]model.Tag}
// @Router       /tag [get]
func (h *tagHandle) GetTags(ctx *fiber.Ctx) error {
	tags, err := h.svc.GetTags()
	if err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	return ctx.JSON(tags)
}

// @Summary      Create a new tag
// @Description  Create a new tag
// @Tags         tag
// @Accept       json
// @Produce      json
// @Param        tag  body      []model.Tag  true  "Tag object"
// @Success      200  {object}  errno.errno
// @Router       /tag [post]
func (h *tagHandle) CreateTag(ctx *fiber.Ctx) error {
	var tag []model.Tag
	if err := ctx.BodyParser(&tag); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(tag, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.CreateTags(tag)
	if err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	return ctx.JSON(errno.OK)
}

// @Summary      Delete a tag by ID
// @Description  Delete a tag by ID
// @Tags         tag
// @Accept       json
// @Produce      json
// @Param        id  body      []uint  true  "Tag IDs"
// @Success      204
// @Router       /tag [delete]
func (h *tagHandle) DeleteTag(ctx *fiber.Ctx) error {
	var body = make([]uint, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.DeleteTags(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      Update a tag by ID
// @Description  Update a tag by ID
// @Tags         tag
// @Accept       json
// @Produce      json
// @Param        id   path      string        true  "Tag ID"
// @Param        tag  body      model.Tag     true  "Tag object"
// @Success      200  {object}  model.Tag
// @Router       /tag/{id} [put]
func (h *tagHandle) UpdateTag(ctx *fiber.Ctx) error {
	// id := ctx.Params("id")
	// if id == "" {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(errno.InvalidIDError())
	// }

	// var tag model.Tag
	// if err := ctx.BodyParser(&tag); err != nil {
	// 	h.logger.Errorf(err.Error())
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(errno.DecodeError(err))
	// }

	// if err := h.valid.Struct(tag); err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(errno.DecodeError(err))
	// }

	// tag.ID = id
	// err := h.svc.UpdateTag(&tag)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(errno.DecodeError(err))
	// }

	return ctx.JSON(errno.OK)
}
