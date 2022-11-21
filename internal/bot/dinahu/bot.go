package dinahu

import (
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/register"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

		switch update.Message.Command() {
		case "register":
			msg.ReplyToMessageID = update.Message.MessageID
			msg.Text, err = register.Handler(update)
			if err != nil {
				msg.Text = err.Error()
			}
		case "delete":
			// TODO : удалить юзера с бд
			msg.Text = "ди наху отсюда ты удален"
		case "start", "run":
			// TODO: выбрать рандомного юзера и послать его наху
			msg.Text = "Да ди ты наху"
		case "stats":
			// TODO: статистику посланных наху юзеров
			msg.Text = "тут будет статистика ди наху пон"
		case "help":
			msg.Text = "Ди наху со своим /help"
		default:
			msg.Text = "Не пон тя попробуй /help"
		}

		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
