package init

import (
	"log"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/bot/dinahu"
)

func Bot() {
	log.Println("Инитиализирую бота наху")
	dinahu.GetBot()
	log.Println("Инитиализировал бота наху")
}
