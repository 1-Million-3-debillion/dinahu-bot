package remove

import (
	"context"
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

const (
	deleteUserFailed string = "не удалось удалить пользователя: %v"
)

func Handler(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "ну и ди наху отсюда.")
	msg.ReplyToMessageID = update.Message.MessageID

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	modelUser, err := user.GetByID(ctx, update.Message.From.ID)
	if err != nil {
		msg.Text = "ты не зарегистрирован ди наху"
		return msg
	}

	tx, err := sqlite.SerializeTransaction(ctx)
	if err != nil {
		msg.Text = err.Error()
		return msg
	}

	if err = modelUser.Delete(ctx, tx); err != nil {
		msg.Text = fmt.Sprintf(deleteUserFailed, err)
		return msg
	}

	if err = tx.Commit(); err != nil {
		msg.Text = err.Error()
		return msg
	}

	return msg
}
