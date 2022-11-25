package dayInfo

import (
	"context"
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/info"
	"github.com/1-Million-3-debillion/dinahu-bot/tools"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

func Handler(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	now := time.Now()

	location, err := time.LoadLocation(tools.Location)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return msg, err
	}

	from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	to := from.Add(24 * time.Hour)

	model, err := info.GetInfo(ctx, from.Unix(), to.Unix())
	if err != nil {
		msg.Text = admin.ErrorMessage
		return msg, err
	}

	msg.Text = fmt.Sprintf(
		"За сегодня %s зарегистрировалось %v чатов и %v пользователей",
		from.Format(tools.DayLayout),
		model.Chats,
		model.Users,
	)

	return msg, nil
}
