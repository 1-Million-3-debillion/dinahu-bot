package postgres

import (
	"log"
	"sync"

	"github.com/1-Million-3-debillion/dinahu-bot/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

const failGetDb string = "GetDB() failed: %v\n"

var (
	db     *sqlx.DB
	onceDb sync.Once
)

func GetDB() *sqlx.DB {
	onceDb.Do(func() {
		var err error
		cfg := config.GetConfig()

		dsn := "postgres://" + cfg.DbUser + ":" + cfg.DbPassword + "@" + cfg.DbHost + cfg.DbPort + "/" + cfg.DbName + "?sslmode=disable"

		db, err = sqlx.Open("pgx", dsn)
		if err != nil {
			log.Fatalf(failGetDb, err)
		}

		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}

		db.SetMaxOpenConns(5)
	})

	return db
}
