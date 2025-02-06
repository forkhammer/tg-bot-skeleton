package api

import (
	"main/bot"
	"main/config"

	"github.com/gin-gonic/gin"
)

type Application struct {
	engine *gin.Engine
	router *Router
}

func NewApplication(bot *bot.Application) *Application {
	engine := gin.Default()
	router := NewRouter(bot)
	return &Application{
		engine: engine,
		router: router,
	}
}

func (a *Application) Run() {
	a.engine.Run(config.Settings.GetHostPort())
}
