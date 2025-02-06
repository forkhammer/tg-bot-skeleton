package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	START_COMMAND = "start"
	HELP_COMMAND  = "help"
)

func getBotCommands(chatId int64) tgbotapi.SetMyCommandsConfig {
	return tgbotapi.NewSetMyCommandsWithScope(
		tgbotapi.NewBotCommandScopeChat(chatId),
		tgbotapi.BotCommand{
			Command:     "/" + START_COMMAND,
			Description: "Регистрация",
		},
		tgbotapi.BotCommand{
			Command:     "/" + HELP_COMMAND,
			Description: "Справка",
		},
	)
}
