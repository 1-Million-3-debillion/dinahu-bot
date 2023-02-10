package chats

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
	msg.Text += "Список чатов:\n\n"

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data, err := chat.GetAll(ctx)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return err
	}

	location, err := time.LoadLocation(tools.Location)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return err
	}

	for _, v := range data {
		msg.Text += fmt.Sprintf(
			"chat_id: %v\nchat_name: %s\nregistered: %v\ncreated_at: %v\n\n",
			v.ChatID,
			v.Name,
			v.Registered,
			v.CreatedAt.In(location).Format(tools.TimeLayout),
		)
	}

	msg.Text += fmt.Sprintf("chats: %v", len(data))

	return nil
}
