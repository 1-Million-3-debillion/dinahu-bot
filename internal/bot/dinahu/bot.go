package dinahu

import (
	"errors"
	"fmt"

	"github.com/1-Million-3-debillion/dinahu-bot/config"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/register"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/remove"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/sendnahu"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/stats"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const help string = `/register - Зарегистрироваться

/sendnahu - Послать рандомного пользователя

/stats - Статистика посланных пользователей

/delete - Удалиться и бот больше не будет вас посылать

/help - Хз вообще что это

Все ди наху`

func Run(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) error {
	var err error

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Chat.ID {
		case config.GetConfig().MillionDebillion:
			msg, err = handleAdminMessages(update)
		default:
			msg, err = handleUserMessages(update)
		}

		if err != nil {
			text := fmt.Sprintf("chat_id: %v\nchat_name: %s\nuser_id: %v\nusername: %s\ncommand: %s\nerror: %s",
				update.Message.Chat.ID,
				update.Message.Chat.Title,
				update.Message.From.ID,
				update.Message.From.UserName,
				update.Message.Text,
				err.Error(),
			)

			errMsg := tgbotapi.NewMessage(config.GetConfig().MillionDebillion, text)
			_, err = bot.Send(errMsg)
			if err != nil {
				return err
			}
		}

		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func handleUserMessages(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	var err error

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.DisableNotification = true

	switch update.Message.Command() {
	case "register":
		msg, err = register.Handler(update)
	case "delete":
		msg, err = remove.Handler(update)
	case "start", "run":
		msg.Text = "Да ди ты наху"
	case "stats":
		msg, err = stats.Handler(update)
	case "help":
		msg.Text = help
	case "sendnahu":
		msg, err = sendnahu.Handler(update)
	case "errortest":
		msg.Text = "сука"
		err = errors.New("ошибка наху")
	default:
		msg.ReplyToMessageID = update.Message.MessageID
		msg.Text = "Ты понимаешь что я не понимаю? попробуй /help"
	}

	return msg, err
}

func handleAdminMessages(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	var err error

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.DisableNotification = true

	switch update.Message.Command() {
	case "chats":
		msg.Text = "тут будет список всех чатов"
	case "users":
		msg.Text = "тут будет список всех юзеров по chat_id"
	default:
		msg.ReplyToMessageID = update.Message.MessageID
		msg.Text = "Я хуй его знает"
	}

	return msg, err
}
