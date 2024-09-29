package di

import (
	"github.com/gsxhnd/garage/garage_server/router"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	r router.Router
}

func NewApplication(r router.Router) *Application {
	return &Application{
		r: r,
	}
}

func (a *Application) Run() error {
	var g errgroup.Group

	g.Go(func() error {
		return a.r.Run()
	})

	if err := g.Wait(); err != nil {
		return err
	} else {
		return nil
	}
}