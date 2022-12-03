package sendnahu

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/handler"

	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/stats"
	"github.com/1-Million-3-debillion/dinahu-bot/internal/storage/sqlite/repo/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var variants = []string{
	"я б вас послал, да вижу — вы оттуда!",
	"ди наху",
	"кет наху",
	"ди наху пон",
	"иди нахуй",
	"хуй на иди",
	"наху ди",
	"Что делать, если вас послали Нахуй?\n\n" +
		"Всерьёз и надолго обидеться на своего собеседника, занеся его в личный чёрный список.\n\n\n",
	"вам билет на пешее эротическое путешествие",
	"ушел в мир биологии пон",
	"на три веселые буквы ди пон",
	"ди подумай",
}

func Handler(update tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data, err := user.GetByChatID(ctx, update.Message.Chat.ID)
	if err != nil {
		msg.Text = handler.ErrorMessage
		return err
	}

	if len(data) <= 1 {
		msg.Text = "Должно быть зарегистрированно 2+ человек /register"
		return nil
	}

	modelUser := data[rand.Intn(len(data))]

	modelStats := stats.Stats{
		UserID: modelUser.UserID,
		ChatID: update.Message.Chat.ID,
	}

	tx, err := sqlite.SerializeTransaction(ctx)
	if err != nil {
		msg.Text = handler.ErrorMessage
		return err
	}

	if err = modelStats.Update(ctx, tx); err != nil {
		msg.Text = handler.ErrorMessage
		return err
	}

	if err = tx.Commit(); err != nil {
		msg.Text = handler.ErrorMessage
		return err
	}

	msg.Text = fmt.Sprintf("@%s %s", modelUser.Username, variants[rand.Intn(len(variants))])

	return nil
}
