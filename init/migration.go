package init

import (
	"log"
	"os"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/postgres"
	"github.com/1-Million-3-debillion/dinahu-bot/tools"
)

const fail string = "%s failed: %v\n"

func Migration(dir string) {
	log.Println("Инитиализирую миграцию наху")

	files, err := tools.GetFiles(dir)
	if err != nil {
		log.Fatal(err)
	}

	db := postgres.GetDB()

	for _, file := range files {
		path := dir + file

		query, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf(fail, file, err)
		}

		_, err = db.Exec(string(query))
		if err != nil {
			log.Fatalf(fail, file, err)
		}
	}

	log.Println("Инитиализировал миграцию наху")
}
