package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

const cfgPath string = "config.json"

type Config struct {
	DinahuToken string `json:"dinahu_token"`
	DbName      string `json:"db_name"`
}

var (
	cfg     Config
	onceCfg sync.Once
)

func GetConfig() *Config {
	onceCfg.Do(func() {
		data, err := os.ReadFile(cfgPath)
		if err != nil {
			log.Fatal(err)
		}

		if err = json.Unmarshal(data, &cfg); err != nil {
			log.Fatal(err)
		}
	})

	return &cfg
}
