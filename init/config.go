package init

import (
	"log"

	"github.com/1-Million-3-debillion/dinahu-bot/config"
)

func Config() {
	log.Println("Инитиализирую конфиг наху")
	config.GetConfig()
	log.Println("Инитиализировал конфиг наху")
}
