package api

import (
	"main/bot"
)

type Router struct {
	bot *bot.Application
}

func NewRouter(bot *bot.Application) *Router {
	return &Router{
		bot: bot,
	}
}
