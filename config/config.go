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
	DbName           string `json:"db_name"`
	MillionDebillion int64  `json:"million_debillion"`
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
