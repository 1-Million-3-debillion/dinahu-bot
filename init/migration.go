package init

import (
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	"github.com/1-Million-3-debillion/dinahu-bot/tools"
	"log"
	"os"
)

const failMigration string = "%s failed: %v"

func Migration(dir string) {
	log.Println("initialize migration")

	files, err := tools.GetFiles(dir)
	if err != nil {
		log.Fatal(err)
	}

	db := sqlite.GetDB()

	for _, file := range files {
		query, err := os.ReadFile(dir + file)
		if err != nil {
			log.Fatalf(failMigration, file, err)
		}

		log.Println(file)

		_, err = db.Exec(string(query))
		if err != nil {
			log.Fatalf(failMigration, file, err)
		}
	}

	log.Println("initialized migration")
}
