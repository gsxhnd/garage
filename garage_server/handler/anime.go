package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/service"
)

type AnimeHandler interface {
	GetAnime(ctx *fiber.Ctx) error
	GetAnimes(ctx *fiber.Ctx) error
	CreateAnime(ctx *fiber.Ctx) error
	DeleteAnime(ctx *fiber.Ctx) error
	UpdateAnime(ctx *fiber.Ctx) error
}

type animeHandle struct {
	valid *validator.Validate
	svc   service.AnimeService
}

func NewAnimeHandler(svc service.AnimeService, v *validator.Validate) AnimeHandler {
	return &animeHandle{
		valid: v,
		svc:   svc,
	}
}

// CreateAnime implements AnimeHandler.
func (a *animeHandle) CreateAnime(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteAnime implements AnimeHandler.
func (a *animeHandle) DeleteAnime(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAnime implements AnimeHandler.
func (a *animeHandle) GetAnime(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAnimes implements AnimeHandler.
func (a *animeHandle) GetAnimes(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateAnime implements AnimeHandler.
func (a *animeHandle) UpdateAnime(ctx *fiber.Ctx) error {
	panic("unimplemented")
}
