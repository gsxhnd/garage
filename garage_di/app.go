package garage_di

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
	garage_ui "github.com/gsxhnd/garage/garage-ui"
	"github.com/gsxhnd/garage/garage_server/routes"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	router *routes.Routes
}

// Open calls the OS default program for uri
func Open(uri string) error {
	var commands = map[string]string{
		"windows": "cmd /c start",
		"darwin":  "open",
		"linux":   "xdg-open",
	}

	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}

func NewApplication(r *routes.Routes) *Application {
	r.Init()
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

	// Open("http://localhost:8081")

	if err := g.Wait(); err != nil {
		return err
	} else {
		return nil
	}
}
