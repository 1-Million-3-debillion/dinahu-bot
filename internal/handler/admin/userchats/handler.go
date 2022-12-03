package userchats

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/chat"
	"github.com/1-Million-3-debillion/dinahu-bot/tools"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handler(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	arr := strings.Split(update.Message.Text, " ")
	if len(arr) < 2 {
		msg.Text = "Неверно введена команда\nПопробуйте так: /userchats user_id/username"
		return nil
	}

	var value string
	for _, v := range arr[1:] {
		value += v + " "
	}

	value = value[:len(value)-1]
	msg.Text = fmt.Sprintf("Чаты где зарегистрирован пользователь %s:\n\n", value)

	data, err := chat.GetByUser(ctx, value)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return err
	}

	if len(data) == 0 {
		msg.Text = "Не нашел ниче босс"
		return nil
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
			time.Unix(v.CreatedAt, 0).In(location).Format(tools.TimeLayout),
		)
	}

	return nil
}
