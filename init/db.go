package init

import (
	"log"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/postgres"
)

func DB() {
	// TODO: перейти на PostgreSQL
	log.Println("Инициализирую бдшку наху")
	// sqlite.GetDB()
	postgres.GetDB()
	log.Println("Инициализировал бдшку наху")
}
