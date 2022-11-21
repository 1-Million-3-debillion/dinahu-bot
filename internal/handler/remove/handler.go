package remove

import (
	"context"
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/userChat"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

const (
	deleteUserFailed string = "не удалось удалить пользователя: %v"
	dbQueryFailed    string = "ошибка на стороне базы: %v"
)

func Handler(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "ну и ди наху отсюда.")
	msg.ReplyToMessageID = update.Message.MessageID

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	has, err := userChat.HasUserInChat(ctx, update.Message.From.ID, update.Message.Chat.ID)
	if err != nil {
		msg.Text = fmt.Sprintf(dbQueryFailed, err)
		return msg
	}

	if !has {
		msg.Text = "ты не зареган ди наху"
		return msg
	}

	modelUserChat := userChat.UserChat{
		UserID: update.Message.From.ID,
		ChatID: update.Message.Chat.ID,
	}

	tx, err := sqlite.SerializeTransaction(ctx)
	if err != nil {
		msg.Text = err.Error()
		return msg
	}

	if err = modelUserChat.DeleteUserFromChat(ctx, tx); err != nil {
		msg.Text = fmt.Sprintf(deleteUserFailed, err)
		return msg
	}

	if err = tx.Commit(); err != nil {
		msg.Text = err.Error()
		return msg
	}

	return msg
}
