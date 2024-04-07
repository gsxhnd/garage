package garage_di

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/garage/utils"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	app *fiber.App
}

func NewApplication(cfg *utils.Config, app *fiber.App) *Application {
	return &Application{
		app: app,
	}
}

func (a *Application) Run() error {
	var g errgroup.Group

	g.Go(func() error {
		return a.app.Listen(":8080")
	})

	if err := g.Wait(); err != nil {
		return err
	} else {
		return nil
	}
}
