package dinahu

import (
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
		msg.DisableNotification = true

		switch update.Message.Command() {
		case "register":
			msg = register.Handler(update)
		case "delete":
			msg = remove.Handler(update)
		case "start", "run":
			msg.Text = "Да ди ты наху"
		case "stats":
			msg = stats.Handler(update)
		case "help":
			msg.Text = help
		case "sendnahu":
			msg = sendnahu.Handler(update)
		default:
			msg.ReplyToMessageID = update.Message.MessageID
			msg.Text = "Ты понимаешь что я не понимаю? попробуй /help"
		}

		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
