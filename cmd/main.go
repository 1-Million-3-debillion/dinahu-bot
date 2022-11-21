package main

import (
	"github.com/1-Million-3-debillion/dinahu-bot/config"
	initialize "github.com/1-Million-3-debillion/dinahu-bot/init"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/bot/dinahu"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func init() {
	config.GetConfig()
	sqlite.GetDB()
	initialize.Migration("./internal/storage/sqlite/migration/")
}

func main() {
	bot, err := tgbotapi.NewBotAPI(config.GetConfig().DinahuToken)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	err = dinahu.Run(bot, updates)
	if err != nil {
		log.Fatal(err)
	}
}
