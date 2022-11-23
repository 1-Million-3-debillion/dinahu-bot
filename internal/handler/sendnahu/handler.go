package sendnahu

import (
	"context"
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler"
	"math/rand"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/stats"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handler(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data, err := user.GetByChatID(ctx, update.Message.Chat.ID)
	if err != nil {
		msg.Text = handler.ErrorMessage
		return msg, err
	}

	if len(data) <= 1 {
		msg.Text = "Должно быть зарегистрированно 2+ человек /register"
		return msg, nil
	}

	modelUser := data[rand.Intn(len(data))]

	modelStats := stats.Stats{
		UserID: modelUser.UserID,
		ChatID: update.Message.Chat.ID,
	}

	tx, err := sqlite.SerializeTransaction(ctx)
	if err != nil {
		msg.Text = handler.ErrorMessage
		return msg, err
	}

	if err = modelStats.Update(ctx, tx); err != nil {
		msg.Text = handler.ErrorMessage
		return msg, err
	}

	if err = tx.Commit(); err != nil {
		msg.Text = handler.ErrorMessage
		return msg, err
	}

	msg.Text = fmt.Sprintf("@%s ди наху", modelUser.Username)

	return msg, nil
}
