package handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/garage_ffmpeg"
	"github.com/gsxhnd/garage/garage_server/task"
	"github.com/gsxhnd/garage/utils"
)

type FFmpegHandler interface {
	Convert(ctx *fiber.Ctx) error
}

type ffmpegHander struct {
	logger      utils.Logger
	validator   *validator.Validate
	taskManager task.TaskMgr
}

func NewFFmpegHandler(l utils.Logger, v *validator.Validate, t task.TaskMgr) FFmpegHandler {
	return &ffmpegHander{
		logger:      l,
		validator:   v,
		taskManager: t,
	}
}

type convertModel struct {
	Name string `json:"name" validate:"required"`
}

func (h *ffmpegHander) Convert(ctx *fiber.Ctx) error {
	body := new(convertModel)

	if err := ctx.BodyParser(body); err != nil {
		h.logger.Errorf("boyd parser error: %s", err.Error())
	}

	if err := h.validator.Struct(body); err != nil {
		h.logger.Errorf("body validation error: %s", err.Error())
	}

	h.logger.Debugf("ffmpeg handler input path: %s", body.Name)

	task, err := task.NewFFmpegTask(&garage_ffmpeg.VideoBatchOption{
		InputPath:    body.Name,
		InputFormat:  "mp4",
		OutputPath:   "/home/gsxhnd/Code/personal/garage/data",
		OutputFormat: "mkv",
		Exec:         true,
	}, "convert")
	if err != nil {
		h.logger.Errorf("init task error: %s", err.Error())
		return nil
	}

	h.taskManager.AddTask(task)
	h.logger.Debugf("Task id: %s", task.GetId())

	for i := range task.GetOB().Observe() {
		fmt.Println(i.V)
	}
	return nil
}
