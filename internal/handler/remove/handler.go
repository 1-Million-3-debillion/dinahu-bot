package remove

import (
	"context"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/postgres"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/postgres/repo/userChat"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handler(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	msg.Text = "ну и ди наху отсюда."

	msg.ReplyToMessageID = update.Message.MessageID

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	has, err := userChat.HasUserInChat(ctx, update.Message.From.ID, update.Message.Chat.ID)
	if err != nil {
		msg.Text = handler.ErrorMessage
		return err
	}

	if !has {
		msg.Text = "ты не зареган ди наху"
		return nil
	}

	modelUserChat := userChat.UserChat{
		UserID: update.Message.From.ID,
		ChatID: update.Message.Chat.ID,
	}

	tx, err := postgres.SerializeTransaction(ctx)
	if err != nil {
		msg.Text = handler.ErrorMessage
		return err
	}

	if err = modelUserChat.DeleteUserFromChat(ctx, tx); err != nil {
		msg.Text = handler.ErrorMessage
		return err
	}

	if err = tx.Commit(); err != nil {
		msg.Text = handler.ErrorMessage
		return err
	}

	return nil
}
