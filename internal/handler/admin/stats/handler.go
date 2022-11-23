package stats

import (
	"context"
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin"
	"strconv"
	"strings"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/stats"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handler(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Статистика посланных пользователей\n")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	arr := strings.Split(update.Message.Text, " ")
	if len(arr) != 2 {
		msg.Text = "Неверно введена команда\nПопробуйте так: /stats chat_id(int64)"
		return msg, nil
	}

	chatID, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		msg.Text = fmt.Sprintf("вторым аргументов должен быть chat_id(int64)")
		return msg, nil
	}

	data, err := stats.GetByChatID(ctx, chatID)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return msg, err
	}

	if len(data) == 0 {
		msg.Text = "Статистика пустая"
		return msg, nil
	}

	for i, v := range data {
		msg.Text += fmt.Sprintf("%v: @%s был послан %v раз\n", i+1, v.Username, v.DinahuCount)
	}

	return msg, nil
}
