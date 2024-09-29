package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/service"
)

type StarHandler interface {
	GetStar(ctx *fiber.Ctx) error
	GetStars(ctx *fiber.Ctx) error
	CreateStars(ctx *fiber.Ctx) error
	DeleteStars(ctx *fiber.Ctx) error
	UpdateStar(ctx *fiber.Ctx) error
}

type starHandle struct {
	// validator *validator.Validate
	// svc service.PingService
}

func NewStarHandler(svc service.PingService) StarHandler {
	return &starHandle{
		// svc: svc,
	}
}

// CreateStars implements StarHandler.
func (s *starHandle) CreateStars(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteStars implements StarHandler.
func (s *starHandle) DeleteStars(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetStar implements StarHandler.
func (s *starHandle) GetStar(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetStars implements StarHandler.
func (s *starHandle) GetStars(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateStar implements StarHandler.
func (s *starHandle) UpdateStar(ctx *fiber.Ctx) error {
	panic("unimplemented")
}
