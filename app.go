package main

import (
	"context"
	"main/api"
	"main/bot"
)

type MainApplication struct {
	bot *bot.Application
	api *api.Application
}

func NewMainApplication() *MainApplication {
	appBot := bot.NewApplication()
	appApi := api.NewApplication(appBot)
	return &MainApplication{
		bot: appBot,
		api: appApi,
	}
}

func (m *MainApplication) Run() {
	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go m.bot.Run(cancelCtx)

	m.api.Run()
}
