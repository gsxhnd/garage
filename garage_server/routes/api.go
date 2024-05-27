package routes

func (r *router) ApiInit() {
	api := r.app.Group("/api")
	api.Use(r.m.RequestLog)
	api.Get("/ping", r.h.RootHandler.Ping)

	// ffmpegGroup := api.Group("/ffmpeg")
	// ffmpegGroup.Post("/convert", r.h.FFmpegHander.Convert)
	// ffmpegGroup.Post("/add_fonts", r.h.FFmpegHander.AddFonts)
	// ffmpegGroup.Post("/add_subtitle", r.h.FFmpegHander.AddSubtitle)

	crawlGroup := api.Group("/crawl")
	crawlGroup.Post("/javbus", r.h.JavHandler.CrawlJavbus)
	crawlGroup.Post("/javdb")
	crawlGroup.Post("/tenhou")

	javGroup := api.Group("/jav")
	javGroup.Get("/")
	javGroup.Get("/code/:code")
	javGroup.Get("/star_code/:code")
	javGroup.Put("/code/:code")
	javGroup.Put("/star_code/:code")
}
