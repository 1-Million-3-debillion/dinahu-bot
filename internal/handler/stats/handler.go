package stats

import (
	"context"
	"fmt"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/stats"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handler(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	msg.Text = "Статистика посланных пользователей\n"

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data, err := stats.GetByChatID(ctx, update.Message.Chat.ID)
	if err != nil {
		msg.Text = handler.ErrorMessage
		return err
	}

	if len(data) == 0 {
		msg.Text = "Статистика пустая епт /sendnahu"
		return nil
	}

	for i, v := range data {
		if v.Username != "" {
			msg.Text += fmt.Sprintf("%v: @%s был послан %v раз\n", i+1, v.Username, v.DinahuCount)
			continue
		}
		if v.FirstName != "" && v.LastName != "" {
			msg.Text += fmt.Sprintf("%v: %s %s был послан %v раз\n", i+1, v.FirstName, v.LastName, v.DinahuCount)
			continue
		}
		if v.FirstName != "" {
			msg.Text += fmt.Sprintf("%v: %s был послан %v раз\n", i+1, v.FirstName, v.DinahuCount)
			continue
		}
		if v.LastName != "" {
			msg.Text += fmt.Sprintf("%v: %s был послан %v раз\n", i+1, v.LastName, v.DinahuCount)
			continue
		}
	}

	return nil
}
