package chatsInfo

import (
	"context"
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/chat"
	"github.com/1-Million-3-debillion/dinahu-bot/tools"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

func Handler(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	loc, err := time.LoadLocation(tools.Location)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return msg, err
	}

	now := time.Now().In(loc)
	from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	to := from.Add(24 * time.Hour)

	data, err := chat.GetByPeriod(ctx, from.Unix(), to.Unix())
	if err != nil {
		msg.Text = admin.ErrorMessage
		return msg, err
	}

	msg.Text += fmt.Sprintf("За сегодня %s зарегистрировалось %v чатов\n\n", now.Format(tools.DayLayout), len(data))

	for _, v := range data {
		msg.Text += fmt.Sprintf(
			"chat_id: %v\nname: %s\ncreated_at: %s\n\n",
			v.ChatID,
			v.Name,
			v.CreatedAt,
			time.Unix(v.CreatedAt, 0).Format(tools.TimeLayout),
		)
	}

	return msg, nil
}
