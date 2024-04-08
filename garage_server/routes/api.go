package routes

func (r *router) ApiInit() {
	api := r.app.Group("/api")
	api.Use(r.m.RequestLog)
	api.Get("/ping", r.h.RootHandler.Ping)

	ffmpegGroup := api.Group("/ffmpeg")
	ffmpegGroup.Get("/videos", r.h.RootHandler.Ping)
}
