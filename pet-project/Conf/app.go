package Conf

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	return &App{
		Router: gin.Default(),
	}
}

func (a *App) Run() err {
	return a.Router.Run(":" + cfg.Port)
}
