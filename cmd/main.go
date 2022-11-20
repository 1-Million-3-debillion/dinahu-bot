package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Config struct {
	DinahuToken string `json:"dinahu_token"`
}

var (
	cfg     Config
	onceCfg sync.Once
)

func GetConfig() *Config {
	onceCfg.Do(func() {
		data, err := os.ReadFile("config.json")
		if err != nil {
			log.Fatal(err)
		}

		if err = json.Unmarshal(data, &cfg); err != nil {
			log.Fatal(err)
		}
	})

	return &cfg
}

func init() {
	GetConfig()
}

func main() {
}
