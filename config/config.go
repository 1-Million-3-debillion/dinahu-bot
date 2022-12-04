package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

const cfgPath string = "config.json"

type Config struct {
	DinahuToken      string `json:"dinahu_token"`
	MillionDebillion int64  `json:"million_debillion"`
	DbUser           string `json:"db_user"`
	DbPassword       string `json:"db_password"`
	DbHost           string `json:"db_host"`
	DbPort           string `json:"db_port"`
	DbName           string `json:"db_name"`
}

var (
	cfg     Config
	onceCfg sync.Once
	fail    string = "GetConfig() failed: %v\n"
)

func GetConfig() *Config {
	onceCfg.Do(func() {
		data, err := os.ReadFile(cfgPath)
		if err != nil {
			log.Fatalf(fail, err)
		}

		if err = json.Unmarshal(data, &cfg); err != nil {
			log.Fatalf(fail, err)
		}
	})

	return &cfg
}
