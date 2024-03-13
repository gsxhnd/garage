package garage_di

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	garage_ui "github.com/gsxhnd/garage/garage-ui"
	"github.com/gsxhnd/garage/garage_server/routes"
	"github.com/gsxhnd/garage/utils"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	router *routes.Routers
}

func NewApplication(cfg *utils.Config, r *routes.Routers) *Application {
	return &Application{
		router: r,
	}
}

func (a *Application) Run() error {
	var g errgroup.Group

	g.Go(func() error {
		var ui = gin.New()
		dist, err := fs.Sub(garage_ui.Web, "dist")
		if err != nil {
			log.Fatalf("dist file server")
			return err
		}

		ui.StaticFS("/", http.FS(dist))
		return ui.Run("0.0.0.0:8081")
	})

	g.Go(func() error {
		return a.router.Engine.Run("0.0.0.0:8080")
	})

	if err := g.Wait(); err != nil {
		return err
	} else {
		return nil
	}
}
