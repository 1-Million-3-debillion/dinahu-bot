package dayInfo

import (
	"context"
	"fmt"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/postgres/repo/info"
	"github.com/1-Million-3-debillion/dinahu-bot/tools"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handler(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	loc, err := time.LoadLocation(tools.Location)
	if err != nil {
		return err
	}

	now := time.Now().In(loc)

	location, err := time.LoadLocation(tools.Location)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return err
	}

	from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	to := from.Add(24 * time.Hour)

	model, err := info.GetInfo(ctx, from, to)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return err
	}

	msg.Text = fmt.Sprintf(
		"За сегодня %s зарегистрировалось %v чатов и %v пользователей",
		from.Format(tools.DayLayout),
		model.Chats,
		model.Users,
	)

	return nil
}
