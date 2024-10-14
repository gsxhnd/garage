// garage_server/handler/tag.go
package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/service"
)

type TagHandler interface {
	GetTag(ctx *fiber.Ctx) error
	GetTags(ctx *fiber.Ctx) error
	CreateTag(ctx *fiber.Ctx) error
	DeleteTag(ctx *fiber.Ctx) error
	UpdateTag(ctx *fiber.Ctx) error
}

type tagHandle struct {
	valid *validator.Validate
	svc   service.TagService
}

func NewTagHandler(svc service.TagService, v *validator.Validate) TagHandler {
	return &tagHandle{
		valid: v,
		svc:   svc,
	}
}

// @Description  Get a single tag by ID
// @Produce      json
// @Success      200 {object} model.Tag
// @Router       /tags/{id} [get]
func (h *tagHandle) GetTag(ctx *fiber.Ctx) error {
	// Implement the logic to get a single tag by ID
	return ctx.Status(200).SendString("GetTag")
}

// @Description  Get all tags
// @Produce      json
// @Success      200 {array} model.Tag
// @Router       /tags [get]
func (h *tagHandle) GetTags(ctx *fiber.Ctx) error {
	// Implement the logic to get all tags
	return ctx.Status(200).SendString("GetTags")
}

// @Description  Create a new tag
// @Produce      json
// @Success      201 {object} model.Tag
// @Router       /tags [post]
func (h *tagHandle) CreateTag(ctx *fiber.Ctx) error {
	// Implement the logic to create a new tag
	return ctx.Status(201).SendString("CreateTag")
}

// @Description  Delete a tag by ID
// @Produce      json
// @Success      204
// @Router       /tags/{id} [delete]
func (h *tagHandle) DeleteTag(ctx *fiber.Ctx) error {
	// Implement the logic to delete a tag by ID
	return ctx.SendStatus(204)
}

// @Description  Update a tag by ID
// @Produce      json
// @Success      200 {object} model.Tag
// @Router       /tags/{id} [put]
func (h *tagHandle) UpdateTag(ctx *fiber.Ctx) error {
	// Implement the logic to update a tag by ID
	return ctx.Status(200).SendString("UpdateTag")
}
