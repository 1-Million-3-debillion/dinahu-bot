package users

import (
	"context"
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/user"
	"github.com/1-Million-3-debillion/dinahu-bot/tools"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"time"
)

func Handler(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	arr := strings.Split(update.Message.Text, " ")
	if len(arr) != 2 {
		msg.Text = "Неверно введена команда\nПопробуйте так: /users chat_id(int64)"
		return msg, nil
	}

	chatID, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		msg.Text = fmt.Sprintf("вторым аргументов должен быть chat_id(int64)")
		return msg, nil
	}

	data, err := user.GetByChatID(ctx, chatID)
	if err != nil {
		msg.Text = admin.ErrorMessage
		return msg, err
	}

	if len(data) == 0 {
		msg.Text = "Чат не найден"
		return msg, nil
	}

	msg.Text += fmt.Sprintf("Список зарегистрированных пользователей в чате %s:\n\n", data[0].ChatName)

	for _, v := range data {
		msg.Text += fmt.Sprintf(
			"user_id: %v\nfirst_name: %s\nlast_name: %s\nusername: @%s\ncreated_at: %v\n\n",
			v.UserID,
			v.FirstName,
			v.LastName,
			v.Username,
			time.Unix(v.CreatedAt, 0).Format(tools.TimeLayout),
		)
	}

	msg.Text += fmt.Sprintf("users: %v", len(data))

	return msg, nil
}
