package bot

import (
	"fmt"
	"main/admin"

	"github.com/getsentry/sentry-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const REGISTER_USER_FLOW = "REGISTER_USER"

type RegisterUserFlow struct {
	bot         *tgbotapi.BotAPI
	state       *State
	adminClient *admin.AdminClient
	userService *UsersService
}

type registerUserStep string

const (
	registerInitStep     registerUserStep = "init"
	registerPasswordStep                  = "password"
)

type RegisterUserData struct {
	Step registerUserStep
}

func NewInitRegisterUserData() *RegisterUserData {
	return &RegisterUserData{
		Step: registerInitStep,
	}
}

func NewRegisterUserFlow(bot *tgbotapi.BotAPI, state *State, adminClient *admin.AdminClient, userService *UsersService) *RegisterUserFlow {
	return &RegisterUserFlow{
		bot:         bot,
		state:       state,
		adminClient: adminClient,
		userService: userService,
	}
}

func (f *RegisterUserFlow) HandleMessage(msg *tgbotapi.Message, data interface{}) error {
	if data == nil {
		return nil
	}

	flowData := data.(RegisterUserData)

	switch flowData.Step {
	case registerInitStep:
		flowData.Step = registerPasswordStep
		f.state.SetState(msg.From.ID, REGISTER_USER_FLOW, flowData)
		response := tgbotapi.NewMessage(msg.Chat.ID, "Введите пароль")
		f.bot.Send(response)
		return nil
	case registerPasswordStep:
		name := fmt.Sprintf("%s %s", msg.From.FirstName, msg.From.LastName)
		user, err := f.userService.RegisterUser(msg.Chat.ID, &name)
		if err != nil {
			sentry.CaptureException(err)
			response := tgbotapi.NewMessage(msg.Chat.ID, err.Error())
			f.bot.Send(response)
			return nil
		}

		f.state.ClearState(msg.From.ID)

		response := tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Пользователь %s зарегистрирован", *(user.Name)))
		f.bot.Send(response)
		f.bot.Request(getBotCommands(msg.Chat.ID))
		return nil
	}
	return nil
}
