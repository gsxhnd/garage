package handler

// type FFmpegHandler interface {
// 	Convert(ctx *fiber.Ctx) error
// 	AddFonts(ctx *fiber.Ctx) error
// 	AddSubtitle(ctx *fiber.Ctx) error
// }

// type ffmpegHander struct {
// 	logger      utils.Logger
// 	validator   *validator.Validate
// 	taskManager task.TaskMgr
// 	svc         service.TaskService
// }

// func NewFFmpegHandler(l utils.Logger, v *validator.Validate, t task.TaskMgr, svc service.TaskService) FFmpegHandler {
// 	return &ffmpegHander{
// 		logger:      l,
// 		validator:   v,
// 		taskManager: t,
// 		svc:         svc,
// 	}
// }

// type model struct {
// 	InputPath   string `json:"inputPath" validate:"required"`
// 	InputFormat string `json:"inputFormat" validate:"required"`
// 	OutputPath  string `json:"outputPath" validate:"required"`
// 	Exec        bool   `json:"exec"`
// }

// type convertModel struct {
// 	model
// 	OutputFormat string `json:"outputFormat" validate:"required"`
// 	Advance      string `json:"advance"`
// }

// type addFontsModel struct {
// 	model
// 	FontsPath string `json:"fontsPath" validate:"required"`
// }

// type addSubtitleModel struct {
// 	model
// 	FontsPath      string `json:"fontsPath"`
// 	InputSubSuffix string `json:"inputSubSuffix" validate:"required"`
// 	InputSubNo     int    `json:"inputSubNo"`
// 	InputSubTitle  string `json:"inputSubTitle" validate:"required"`
// 	InputSubLang   string `json:"inputSublang" validate:"required"`
// }

// func (h *ffmpegHander) Convert(ctx *fiber.Ctx) error {
// 	body := new(model)

// 	if err := ctx.BodyParser(body); err != nil {
// 		h.logger.Errorf("boyd parser error: %s", err.Error())
// 		return nil
// 	}

// 	if err := h.validator.Struct(body); err != nil {
// 		h.logger.Errorf("body validation error: %s", err.Error())
// 		return nil
// 	}

// 	task, err := task.NewFFmpegTask(&garage_ffmpeg.VideoBatchOption{
// 		InputPath:    body.InputPath,
// 		InputFormat:  "mp4",
// 		OutputPath:   "/home/gsxhnd/Code/personal/garage/data",
// 		OutputFormat: "mkv",
// 		Exec:         true,
// 	}, "convert")
// 	if err != nil {
// 		h.logger.Errorf("init task error: %s", err.Error())
// 		return nil
// 	}

// 	h.taskManager.AddTask(task)
// 	h.logger.Debugf("Task id: %s", task.GetId())

// 	for i := range task.GetOB().Observe() {
// 		fmt.Println(i.V)
// 	}
// 	return nil
// }

// func (h *ffmpegHander) AddFonts(ctx *fiber.Ctx) error {
// 	body := new(addFontsModel)

// 	if err := ctx.BodyParser(body); err != nil {
// 		h.logger.Errorf("boyd parser error: %s", err.Error())
// 		return nil
// 	}

// 	if err := h.validator.Struct(body); err != nil {
// 		h.logger.Errorf("body validation error: %s", err.Error())
// 		return nil
// 	}

// 	task, err := task.NewFFmpegTask(&garage_ffmpeg.VideoBatchOption{
// 		InputPath:    body.InputPath,
// 		InputFormat:  body.InputFormat,
// 		OutputPath:   body.OutputPath,
// 		OutputFormat: body.InputFormat,
// 		FontsPath:    body.FontsPath,
// 		Exec:         body.Exec,
// 	}, "add_fonts")
// 	if err != nil {
// 		h.logger.Errorf("init task error: %s", err.Error())
// 		return nil
// 	}

// 	h.taskManager.AddTask(task)
// 	h.logger.Debugf("Task id: %s", task.GetId())

// 	var data = map[string]interface{}{
// 		"id":   task.GetId(),
// 		"cmds": task.GetCmds(),
// 	}

// 	return ctx.JSON(data)
// }

// func (h *ffmpegHander) AddSubtitle(ctx *fiber.Ctx) error {
// 	body := new(addSubtitleModel)

// 	if err := ctx.BodyParser(body); err != nil {
// 		h.logger.Errorf("boyd parser error: %s", err.Error())
// 		return nil
// 	}

// 	if err := h.validator.Struct(body); err != nil {
// 		h.logger.Errorf("body validation error: %s", err.Error())
// 		return nil
// 	}

// 	task, err := task.NewFFmpegTask(&garage_ffmpeg.VideoBatchOption{
// 		InputPath:      body.InputPath,
// 		InputFormat:    body.InputFormat,
// 		OutputPath:     body.OutputPath,
// 		OutputFormat:   body.InputFormat,
// 		FontsPath:      body.FontsPath,
// 		InputSubSuffix: body.InputSubSuffix,
// 		InputSubNo:     body.InputSubNo,
// 		InputSubTitle:  body.InputSubTitle,
// 		InputSubLang:   body.InputSubLang,
// 		Exec:           body.Exec,
// 	}, "add_subtitle")
// 	if err != nil {
// 		h.logger.Errorf("init task error: %s", err.Error())
// 		return nil
// 	}

// 	h.taskManager.AddTask(task)
// 	h.logger.Debugf("Task id: %s", task.GetId())

// 	var data = map[string]interface{}{
// 		"id":   task.GetId(),
// 		"cmds": task.GetCmds(),
// 	}

// 	return ctx.JSON(data)
// }
