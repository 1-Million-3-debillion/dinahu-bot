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
)

func GetBot() *tgbotapi.BotAPI {
	botOnce.Do(func() {
		var err error
		bot, err = tgbotapi.NewBotAPI(config.GetConfig().DinahuToken)
		if err != nil {
			log.Fatal(err)
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

		msg := tgbotapi.MessageConfig{}

		switch update.Message.Chat.ID {
		case config.GetConfig().MillionDebillion:
			msg, err = handleAdminMessages(update)
		default:
			msg, err = handleUserMessages(update)
		}

		if err != nil {
			sendErrToAdmins(update, err)
		}

		_, err = GetBot().Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
