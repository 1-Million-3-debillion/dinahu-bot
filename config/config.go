package config

import (
	"log"
	"sync"

	"github.com/1-Million-3-debillion/dinahu-bot/tools/env"
	"github.com/joho/godotenv"
)

const cfgPath string = ".env"

type Config struct {
	DinahuToken      string `env:"DINAHU_TOKEN"`
	MillionDebillion int64  `env:"MILLION_DEBILLION"`
	Postgres         struct {
		User     string `env:"POSTGRES_USER"`
		Password string `env:"POSTGRES_PASSWORD"`
		Host     string `env:"POSTGRES_HOST"`
		Port     string `env:"POSTGRES_PORT"`
		Name     string `env:"POSTGRES_DB"`
	}
}

var (
	cfg     = &Config{}
	onceCfg sync.Once
	fail    string = "GetConfig() failed: %v\n"
)

func GetConfig() *Config {
	onceCfg.Do(func() {
		var err error
		err = godotenv.Load(cfgPath)
		if err != nil {
			log.Fatal(err)
		}

		err = env.Unmarshal(cfg)
		if err != nil {
			log.Fatal(err)
		}
	})

	return cfg
}
