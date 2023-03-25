package sqlite

import (
	"log"
	"sync"

	"github.com/1-Million-3-debillion/dinahu-bot/config"
	"github.com/jmoiron/sqlx"
)

const fail string = "GetDB() failed: %v\n"

var (
	db     *sqlx.DB
	onceDb sync.Once
)

func GetDB() *sqlx.DB {
	onceDb.Do(func() {
		var err error
		db, err = sqlx.Open("sqlite3", config.GetConfig().Postgres.Name)
		if err != nil {
			log.Fatalf(fail, err)
		}
	})

	return db
}
