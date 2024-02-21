package di

import "github.com/gsxhnd/garage/garage_server/routes"

type Application struct {
	router *routes.Routes
}

func NewApplication(r *routes.Routes) *Application {
	r.Init()
	return &Application{
		router: r,
	}
}

func (a *Application) Run() error {
	return a.router.Engine.Run("0.0.0.0:8080")
}
