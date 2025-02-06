package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CONFIRM_YES                  = "да"
	CONFIRM_NO                   = "нет"
	ORDER_DELIVERY_WITHOUT_PHOTO = "Без фотографий"
)

func GetConfirmKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	keyword := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CONFIRM_YES),
			tgbotapi.NewKeyboardButton(CONFIRM_NO),
		),
	)
	return &keyword
}
