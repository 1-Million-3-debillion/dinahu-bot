package register

import (
	"context"
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/chat"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/user"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/userChat"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	addUserFailed     string = "не удалось зарегистрировать пользователя: %v"
	addChatFailed     string = "не удалось зарегистрировать чат: %v"
	addUserChatFailed string = "не удалось связать юзера с чатом: %v"
)

func Handler(update tgbotapi.Update) (string, error) {
	msg := "Зарегистрировал тебя а теперь ди наху"
	modelUser := user.User{
		UserID:    update.Message.From.ID,
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
		Username:  update.Message.From.UserName,
	}

	modelChat := chat.Chat{
		ChatID: update.Message.Chat.ID,
		Name:   update.Message.Chat.Title,
	}

	modelUserChat := userChat.UserChat{
		ID:     uuid.NewV4().String(),
		UserID: modelUser.UserID,
		ChatID: modelChat.ChatID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := sqlite.SerializeTransaction(ctx)
	if err != nil {
		return "", err
	}

	if err = modelUser.Add(ctx, tx); err != nil {
		return "", fmt.Errorf(addUserFailed, err)
	}

	if update.Message.Chat.Type != "private" {
		if err = modelChat.Add(ctx, tx); err != nil {
			return "", fmt.Errorf(addChatFailed, err)
		}
	}

	has, err := modelUserChat.HasUserInChat(ctx, modelUser.UserID, modelChat.ChatID)
	if err != nil {
		return "", err
	}

	switch has {
	case true:
		msg = "Ты уже зарегистрирован тут ди наху"
	case false:
		if err = modelUserChat.Add(ctx, tx); err != nil {
			return "", fmt.Errorf(addUserChatFailed, err)
		}
	}

	if err = tx.Commit(); err != nil {
		return "", err
	}

	return msg, nil
}
