package routes

func (r *router) ApiInit() {
	api := r.app.Group("/api")
	api.Use(r.m.RequestLog)
	api.Get("/ping", r.h.RootHandler.Ping)

	ffmpegGroup := api.Group("/ffmpeg")
	ffmpegGroup.Post("/convert", r.h.FFmpegHander.Convert)
	ffmpegGroup.Post("/add_fonts", r.h.FFmpegHander.AddFonts)
	ffmpegGroup.Post("/add_subtitle", r.h.FFmpegHander.Convert)

	javGroup := api.Group("/jav")
	javGroup.Post("/code", r.h.JavHandler.CrawlJavByCode)
}
