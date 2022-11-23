package userchats

import (
	"context"
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/chat"
	"github.com/1-Million-3-debillion/dinahu-bot/tools"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"time"
)

func Handler(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Чаты где зарегистрирован пользователь\n")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	arr := strings.Split(update.Message.Text, " ")
	if len(arr) < 2 {
		msg.Text = "Неверно введена команда\nПопробуйте так: /userchats user_id/username"
		return msg, nil
	}

	var value string
	for _, v := range arr[1:] {
		value += v + " "
	}

	data, err := chat.GetByUser(ctx, value[:len(value)-1])
	if err != nil {
		msg.Text = admin.ErrorMessage
		return msg, err
	}

	if len(data) == 0 {
		msg.Text = "Не нашел ниче босс"
		return msg, nil
	}

	location, err := time.LoadLocation(tools.Location)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return msg, err
	}

	for _, v := range data {
		msg.Text += fmt.Sprintf(
			"chat_id: %v\nchat_name: %s\nregistered: %v\ncreated_at: %v\n\n",
			v.ChatID,
			v.Name,
			v.Registered,
			time.Unix(v.CreatedAt, 0).In(location).Format(tools.TimeLayout),
		)
	}

	return msg, nil
}
