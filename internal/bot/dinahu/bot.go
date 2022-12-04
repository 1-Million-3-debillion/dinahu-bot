package dinahu

import (
	"log"
	"sync"

	"github.com/1-Million-3-debillion/dinahu-bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	bot     *tgbotapi.BotAPI
	botOnce sync.Once
	fail    string = "GetBot() failed: %v\n"
)

func GetBot() *tgbotapi.BotAPI {
	botOnce.Do(func() {
		var err error
		bot, err = tgbotapi.NewBotAPI(config.GetConfig().DinahuToken)
		if err != nil {
			log.Fatalf(fail, err)
		}

		log.Printf("Authorized on account %s", bot.Self.UserName)
	})

	return bot
}

func HandleUpdates(updates tgbotapi.UpdatesChannel) error {
	var err error

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		if update.Message.Chat.ID == update.Message.From.ID {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Chat.ID {
		case config.GetConfig().MillionDebillion:
			err = handleAdminMessages(update, &msg)
		default:
			err = handleUserMessages(update, &msg)
		}

		if err != nil {
			log.Println(err)
			sendErrToAdmins(update, err)
		}

		msg.DisableNotification = true

		_, err = GetBot().Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
