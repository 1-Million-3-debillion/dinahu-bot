package stats

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/stats"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handler(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	msg.Text += "Статистика посланных пользователей\n"

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	arr := strings.Split(update.Message.Text, " ")
	if len(arr) != 2 {
		msg.Text = "Неверно введена команда\nПопробуйте так: /stats chat_id(int64)"
		return nil
	}

	chatID, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		msg.Text = fmt.Sprintf("вторым аргументов должен быть chat_id(int64)")
		return nil
	}

	data, err := stats.GetByChatID(ctx, chatID)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return err
	}

	if len(data) == 0 {
		msg.Text = "Статистика пустая"
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
