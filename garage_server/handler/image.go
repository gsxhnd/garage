package handler

import (
	"fmt"
	"mime"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_server/storage"
	"github.com/gsxhnd/garage/utils"
)

type ImageHandler interface {
	GetMovieImage(ctx *fiber.Ctx) error
	GetStarImage(ctx *fiber.Ctx) error
}

type imageHandle struct {
	logger  utils.Logger
	valid   *validator.Validate
	storage storage.Storage
}

func NewImageHandler(v *validator.Validate, s storage.Storage, l utils.Logger) ImageHandler {
	return &imageHandle{
		logger:  l,
		valid:   v,
		storage: s,
	}
}

func (h *imageHandle) GetMovieImage(ctx *fiber.Ctx) error {
	fmt.Println(ctx.Params("id"))
	return nil
}

func (h *imageHandle) GetStarImage(ctx *fiber.Ctx) error {
	// var id = ctx.Params("id")
	b, format, err := h.storage.GetImage("star", 0, "")
	h.logger.Errorw("GET Star Image", "error", err)
	ctx.Response().Header.Set("Content-Type", mime.TypeByExtension("."+format))
	ctx.Write(b)

	return nil
}
