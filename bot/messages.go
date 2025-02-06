package bot

import (
	"log"
	"main/admin"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessagesService struct {
	bot         *tgbotapi.BotAPI
	adminClient *admin.AdminClient
	userService *UsersService
	state       *State
}

func NewMessagesService(bot *tgbotapi.BotAPI, adminClient *admin.AdminClient, userService *UsersService) *MessagesService {
	return &MessagesService{
		bot:         bot,
		adminClient: adminClient,
		userService: userService,
		state:       NewState(),
	}
}

func (s *MessagesService) ProcessMessage(msg *tgbotapi.Message) error {
	log.Printf("[%s] %s", msg.From.UserName, msg.Text)

	if msg.IsCommand() {
		switch msg.Command() {
		case START_COMMAND:
			flow := NewRegisterUserFlow(s.bot, s.state, s.adminClient, s.userService)
			flow.HandleMessage(msg, *NewInitRegisterUserData())
			return nil
		case HELP_COMMAND:
			return s.processHelpCommand(msg)
		}
	}

	switch msg.Text {
	default:
		return s.processUnknownMessage(msg)
	}
}

func (s *MessagesService) ProcessCallback(callback *tgbotapi.CallbackQuery) error {
	split := strings.Split(callback.Data, " ")

	switch split[0] {
	}
	return nil
}

func (s *MessagesService) SendMessageToAllUsers(text string) error {
	users, err := s.userService.GetAllUsers()
	if err != nil {
		return err
	}
	for _, user := range users {
		response := tgbotapi.NewMessage(user.ChatId, text)
		s.bot.Send(response)
	}
	return nil
}

func (s *MessagesService) processUnknownMessage(msg *tgbotapi.Message) error {
	kind, data := s.state.GetState(msg.From.ID)
	switch kind {
	case REGISTER_USER_FLOW:
		flow := NewRegisterUserFlow(s.bot, s.state, s.adminClient, s.userService)
		return flow.HandleMessage(msg, data)
	default:
		return nil
	}
}

func (s *MessagesService) sendUnregisterResponse(msg *tgbotapi.Message) error {
	response := tgbotapi.NewMessage(msg.Chat.ID, "Вы не зарегистрированы")
	_, err := s.bot.Send(response)
	return err
}

func (s *MessagesService) processHelpCommand(msg *tgbotapi.Message) error {
	tpl, err := RenderTemplate(HELP_TPL, nil)

	if err != nil {
		return err
	}

	response := tgbotapi.NewMessage(msg.Chat.ID, tpl)
	s.bot.Send(response)
	return nil
}
