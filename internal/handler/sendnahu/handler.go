package sendnahu

import (
	"context"
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/user"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/userChat"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"time"
)

func Handler(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data, err := user.GetByChatID(ctx, update.Message.Chat.ID)
	if err != nil {
		msg.Text = "Не удалось получить список пользователей"
		return msg
	}

	if len(data) == 0 {
		msg.Text = "Сначало зарегайся епт /register"
		return msg
	}

	if len(data) == 1 {
		msg.Text = "Должно быть зарегистрированно 2+ человек"
		return msg
	}

	modelUser := data[rand.Intn(len(data))]

	modelUserChat := userChat.UserChat{
		UserID: modelUser.UserID,
		ChatID: update.Message.Chat.ID,
	}

	tx, err := sqlite.SerializeTransaction(ctx)
	if err != nil {
		msg.Text = err.Error()
		return msg
	}

	if err = modelUserChat.AddDinahu(ctx, tx); err != nil {
		msg.Text = "Не удалось послать наху"
		return msg
	}

	if err = tx.Commit(); err != nil {
		msg.Text = err.Error()
		return msg
	}

	msg.Text = fmt.Sprintf("@%s ди наху", modelUser.Username)

	return msg
}
