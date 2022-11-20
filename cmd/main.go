package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	DinahuToken string `json:"dinahu_token"`
}

var (
	cfg     Config
	onceCfg sync.Once
)

var (
	db     *sql.DB
	onceDb sync.Once
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

func GetDb() *sql.DB {
	onceDb.Do(func() {
		var err error
		db, err = sql.Open("sqlite3", "dinahu.db")
		if err != nil {
			log.Fatal(err)
		}
	})

	return db
}

func Migration() {
	query := `
		CREATE TABLE IF NOT EXISTS "user" (
			"user_id" INTEGER UNIQUE NOT NULL,
			"first_name" TEXT,
			"last_name" TEXT,
			"username" TEXT
		);
		
		CREATE TABLE IF NOT EXISTS "chat" (
			"chat_id" INTEGER NOT NULL UNIQUE,
			"name" TEXT
		);

		CREATE TABLE IF NOT EXISTS "user_chat" (
			"id" TEXT NOT NULL UNIQUE,
			"user_id" INTEGER NOT NULL,
			"chat_id" INTGER NOT NULL
		);
		`

	_, err := GetDb().Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	UserID    int64
	FirstName string
	LastName  string
	Username  string
}

type Chat struct {
	ChatID int64
	Name   string
}

type UserChat struct {
	ID     string
	UserID int64
	ChatID int64
}

func init() {
	GetConfig()
	GetDb()
	Migration()
}

func main() {
	bot, err := tgbotapi.NewBotAPI(GetConfig().DinahuToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "register":
			// TODO : зарегать юзера наху
			msg.Text = "тут будет регистрация юзеров"
		case "delete":
			// TODO : удалить юзера с бд
			msg.Text = "ди наху отсюда ты удален"
		case "start", "run":
			// TODO: выбрать рандомного юзера и послать его наху
			msg.Text = "Да ди ты наху"
		case "stats":
			// TODO: статистику посланных наху юзеров
			msg.Text = "тут будет статистика ди наху пон"
		case "help":
			msg.Text = "Ди наху со своим /help"
		default:
			continue
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
