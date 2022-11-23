package dinahu

import (
	"context"
	"errors"
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin/chats"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin/dayInfo"
	adminStats "github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin/stats"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin/users"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/register"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/remove"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/sendnahu"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/stats"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/chat"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleUserMessages(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	var err error

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.DisableNotification = true

	switch update.Message.Command() {
	case "register":
		var has bool
		has, err = chat.HasByID(context.Background(), update.Message.Chat.ID)
		if err != nil {
			return msg, err
		}

		if !has {
			sendMsgToAdmins(fmt.Sprintf(
				"Добавлен новый чат:\nchat_id: %v\nchat_name: %s\n#chat",
				update.Message.Chat.ID,
				update.Message.Chat.Title,
			))
		}

		msg, err = register.Handler(update)
	case "delete":
		msg, err = remove.Handler(update)
	case "start", "run":
		msg.Text = "Да ди ты наху"
	case "stats":
		msg, err = stats.Handler(update)
	case "help":
		msg.Text = help
	case "sendnahu", "nahu":
		msg, err = sendnahu.Handler(update)
	case "errortest":
		msg.Text = "сука"
		err = errors.New("ошибка наху")
	default:
		msg.ReplyToMessageID = update.Message.MessageID
		msg.Text = "Ты понимаешь что я не понимаю? попробуй /help"
	}

	return msg, err
}

func handleAdminMessages(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	var err error

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.DisableNotification = true

	switch update.Message.Command() {
	case "chats":
		msg, err = chats.Handler(update)
	case "users":
		msg, err = users.Handler(update)
	case "stats":
		msg, err = adminStats.Handler(update)
	case "dayinfo":
		msg, err = dayInfo.Handler(update)
	case "help":
		msg.ReplyToMessageID = update.Message.MessageID
		msg.Text = adminHelp
	default:
		msg.ReplyToMessageID = update.Message.MessageID
		msg.Text = "Ты втираешь мне какую то дичь"
	}

	return msg, err
}
