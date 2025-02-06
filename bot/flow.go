package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Flow interface {
	HandleMessage(msg *tgbotapi.Message, data interface{}) error
}
