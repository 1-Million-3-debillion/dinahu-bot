package chatsInfo

import (
	"context"
	"fmt"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/postgres/repo/chat"
	"github.com/1-Million-3-debillion/dinahu-bot/tools"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handler(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	loc, err := time.LoadLocation(tools.Location)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return err
	}

	now := time.Now().In(loc)
	from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	to := from.Add(24 * time.Hour)

	data, err := chat.GetByPeriod(ctx, from, to)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return err
	}

	msg.Text += fmt.Sprintf("За сегодня %s зарегистрировалось %v чатов\n\n", now.Format(tools.DayLayout), len(data))

	for _, v := range data {
		msg.Text += fmt.Sprintf(
			"chat_id: %v\nname: %s\ncreated_at: %s\n\n",
			v.ChatID,
			v.Name,
			v.CreatedAt.Format(tools.TimeLayout),
		)
	}

	return nil
}
