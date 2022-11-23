package register

import (
	"context"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/chat"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/stats"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/user"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/userChat"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	uuid "github.com/satori/go.uuid"
)

func Handler(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Зарегистрировал тебя а теперь ди наху")
	msg.ReplyToMessageID = update.Message.MessageID

	modelUser := user.User{
		UserID:    update.Message.From.ID,
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
		Username:  update.Message.From.UserName,
		CreatedAt: time.Now().Unix(),
	}

	modelChat := chat.Chat{
		ChatID:    update.Message.Chat.ID,
		Name:      update.Message.Chat.Title,
		CreatedAt: time.Now().Unix(),
	}

	modelUserChat := userChat.UserChat{
		ID:     uuid.NewV4().String(),
		UserID: modelUser.UserID,
		ChatID: modelChat.ChatID,
	}

	modelStats := stats.Stats{
		ID:          uuid.NewV4().String(),
		UserID:      modelUser.UserID,
		ChatID:      modelChat.ChatID,
		DinahuCount: 0,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := sqlite.SerializeTransaction(ctx)
	if err != nil {
		msg.Text = handler.ErrorMessage
		return msg, err
	}

	if err = modelUser.Add(ctx, tx); err != nil {
		msg.Text = handler.ErrorMessage
		return msg, err
	}

	if update.Message.Chat.ID != update.Message.From.ID {
		if err = modelChat.Add(ctx, tx); err != nil {
			msg.Text = handler.ErrorMessage
			return msg, err
		}

		has, err := userChat.HasUserInChat(ctx, modelUser.UserID, modelChat.ChatID)
		if err != nil {
			_ = tx.Rollback()
			msg.Text = handler.ErrorMessage
			return msg, err
		}

		switch has {
		case true:
			msg.Text = "Ты уже зарегистрирован тут ди наху"
		case false:
			if err = modelUserChat.Add(ctx, tx); err != nil {
				msg.Text = handler.ErrorMessage
				return msg, err
			}

			if err = modelStats.Add(ctx, tx); err != nil {
				msg.Text = handler.ErrorMessage
				return msg, err
			}
		}
	}

	if err = tx.Commit(); err != nil {
		msg.Text = handler.ErrorMessage
		return msg, err
	}

	return msg, nil
}
