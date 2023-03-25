package main

import (
	"fmt"
	"log"

	"github.com/1-Million-3-debillion/dinahu-bot/config"
	initialize "github.com/1-Million-3-debillion/dinahu-bot/init"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/bot/dinahu"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	initialize.Config()
	cfg := config.GetConfig()
	fmt.Printf("%+v\n", cfg)
	initialize.DB()
	initialize.Migration("./internal/storage/postgres/migration/")
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
