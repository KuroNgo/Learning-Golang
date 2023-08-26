package conf

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Route *gin.Engine
}

func NewApp() *App {
	return &App{
		Route: gin.Default(),
	}
}

func (app *App) Run() error {
	return app.Route.Run(":8080")
}
