package init

import (
	"log"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
)

func DB() {
	// TODO: перейти на PostgreSQL
	log.Println("Инициализирую бдшку наху")
	sqlite.GetDB()
	log.Println("Инициализировал бдшку наху")
}
