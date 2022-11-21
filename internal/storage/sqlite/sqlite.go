package sqlite

import (
	"github.com/1-Million-3-debillion/dinahu-bot/config"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

const failGetDB string = "GetDB() failed: %v"

var (
	db     *sqlx.DB
	onceDb sync.Once
)

func GetDB() *sqlx.DB {
	onceDb.Do(func() {
		var err error
		db, err = sqlx.Open("sqlite3", config.GetConfig().DbName)
		if err != nil {
			log.Fatalf(failGetDB, err)
		}
	})

	return db
}
