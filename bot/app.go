package bot

import (
	"context"
	"log"
	"main/admin"
	"main/config"

	"github.com/getsentry/sentry-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Application struct {
	bot             *tgbotapi.BotAPI
	messagesService *MessagesService
	dbConnection    *DbConnection
}

func NewApplication() *Application {
	bot, err := tgbotapi.NewBotAPI(config.Settings.Token)

	if err != nil {
		sentry.CaptureException(err)
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	dbConnection, err := NewDbConnection(config.Settings.Db.Host, config.Settings.Db.Port, config.Settings.Db.DBName, config.Settings.Db.User, config.Settings.Db.Pass)

	if err != nil {
		sentry.CaptureException(err)
		log.Panic(err)
	}

	userService := NewUsersService(dbConnection)
	messagesService := NewMessagesService(bot, admin.NewAdminClient(config.Settings.AdminUrl, config.Settings.AdminToken), userService)

	app := &Application{
		bot:             bot,
		messagesService: messagesService,
		dbConnection:    dbConnection,
	}

	return app
}

func (a *Application) Run(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := a.bot.GetUpdatesChan(u)

	for {
		select {
		case update := <-updates:
			if update.Message != nil {
				err := a.messagesService.ProcessMessage(update.Message)

				if err != nil {
					sentry.CaptureException(err)
				}
			} else if update.CallbackQuery != nil {
				err := a.messagesService.ProcessCallback(update.CallbackQuery)

				if err != nil {
					sentry.CaptureException(err)
				}
			}
		case <-ctx.Done():
			log.Println("Bot exit")
			return
		}
	}
}
