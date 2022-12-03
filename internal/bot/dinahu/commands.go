package dinahu

import (
	"context"
	"errors"
	"fmt"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin/chats"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin/chatsInfo"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin/dayInfo"
	adminStats "github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin/stats"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin/userchats"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin/users"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/admin/usersInfo"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/register"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/remove"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/sendnahu"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler/stats"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/chat"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleUserMessages(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	var err error

	switch update.Message.Command() {
	case "register":
		var has bool
		has, err = chat.HasByID(context.Background(), update.Message.Chat.ID)
		if err != nil {
			return err
		}

		if !has {
			sendMsgToAdmins(fmt.Sprintf(
				"Добавлен новый чат:\nchat_id: %v\nchat_name: %s\n#chat",
				update.Message.Chat.ID,
				update.Message.Chat.Title,
			))
		}

		err = register.Handler(update, msg)
	case "delete":
		err = remove.Handler(update, msg)
	case "start", "run":
		msg.Text = "Да ди ты наху"
	case "stats":
		err = stats.Handler(update, msg)
	case "help":
		msg.Text = help
	case "sendnahu", "nahu":
		err = sendnahu.Handler(update, msg)
	case "errortest":
		msg.Text = "Ты понимаешь что я не понимаю? попробуй /help"
		err = errors.New("ошибка наху")
	default:
		msg.ReplyToMessageID = update.Message.MessageID
		msg.Text = "Ты понимаешь что я не понимаю? попробуй /help"
	}

	return err
}

func handleAdminMessages(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	var err error

	switch update.Message.Command() {
	case "chats":
		err = chats.Handler(update, msg)
	case "userchats":
		err = userchats.Handler(update, msg)
	case "users":
		err = users.Handler(update, msg)
	case "stats":
		err = adminStats.Handler(update, msg)
	case "dayinfo":
		err = dayInfo.Handler(update, msg)
	case "chatsinfo":
		err = chatsInfo.Handler(update, msg)
	case "usersinfo":
		err = usersInfo.Handler(update, msg)
	case "help":
		msg.ReplyToMessageID = update.Message.MessageID
		msg.Text = adminHelp
	default:
		msg.ReplyToMessageID = update.Message.MessageID
		msg.Text = "Ты втираешь мне какую то дичь"
	}

	return err
}
