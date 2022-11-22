package stats

import (
	"context"
	"fmt"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/stats"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	errorMessage string = "Что то пошло не так. Админы скоро пофиксят"
)

func Handler(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Статистика посланных пользователей\n")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data, err := stats.GetByChatID(ctx, update.Message.Chat.ID)
	if err != nil {
		msg.Text = errorMessage
		return msg, err
	}

	if len(data) == 0 {
		msg.Text = "Статистика пустая епт /sendnahu"
		return msg, nil
	}

	for i, v := range data {
		msg.Text += fmt.Sprintf("%v: @%s был послан %v раз\n", i+1, v.Username, v.DinahuCount)
	}

	return msg, nil
}
