package app

import (
	"Duna/app/errors"
	"Duna/database/repositories"
	"github.com/gin-gonic/gin"
)

type App struct {
	Repositories *repositories.Repositories
	Router       *gin.Engine
}

func NewApp(repositories *repositories.Repositories, router *gin.Engine) *App {
	return &App{
		Repositories: repositories,
		Router:       router,
	}
}

func (app App) Start() error {
	r := app.Router.Run()

	if r != nil {
		return errors.AppStartError{Err: r}
	}

	return nil
}
