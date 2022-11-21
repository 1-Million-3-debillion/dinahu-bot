package stat

import (
	"context"
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/userChat"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

func Handler(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Статистика посланных пользователей\n")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data, err := userChat.GetByChatID(ctx, update.Message.Chat.ID)
	if err != nil {
		msg.Text = fmt.Sprintf("не удалось получить статистику: %v", err)
		return msg
	}

	for i, v := range data {
		msg.Text += fmt.Sprintf("%v: @%s был послан %v раз\n", i+1, v.Username, v.DinahuCount)
	}

	return msg
}
