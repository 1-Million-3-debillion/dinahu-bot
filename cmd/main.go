package main

import (
	"log"

	initialize "github.com/1-Million-3-debillion/dinahu-bot/init"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/bot/dinahu"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	initialize.Config()
	initialize.DB()
	initialize.Migration("./internal/storage/sqlite/migration/")
	initialize.Bot()
}

func main() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := dinahu.GetBot().GetUpdatesChan(u)

	err := dinahu.HandleUpdates(updates)
	if err != nil {
		log.Fatalf("HandleUpdates() failed: %v\n", err)
	}
}
