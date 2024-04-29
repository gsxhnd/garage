package routes

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
	garage_ui "github.com/gsxhnd/garage/garage-ui"
)

func (r *router) EnableWeb() error {
	if !r.cfg.WebEnable {
		return nil
	}

	dist, err := fs.Sub(garage_ui.Web, "dist")
	if err != nil {
		log.Fatalf("dist file server")
		return err
	}

	webG := r.app.Group("/")
	webG.All("*", filesystem.New(filesystem.Config{
		Root:  http.FS(dist),
		Index: "index.html",
	}))
	return nil
}
