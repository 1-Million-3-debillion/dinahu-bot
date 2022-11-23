package dinahu

import (
	"fmt"
	"github.com/1-Million-3-debillion/dinahu-bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const errorMsg string = "МУЖЫКИ Я НЕ СПРАВИЛСЯ ПОМОГИТЕ\nchat_id: %v\nchat_name: %s\nuser_id: %v\nusername: @%s\ncommand: %s\nerror: %s\n#error"

func sendMsgToAdmins(text string) {
	msg := tgbotapi.NewMessage(config.GetConfig().MillionDebillion, text)

	_, err := GetBot().Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func sendErrToAdmins(update tgbotapi.Update, err error) {
	text := fmt.Sprintf(errorMsg,
		update.Message.Chat.ID,
		update.Message.Chat.Title,
		update.Message.From.ID,
		update.Message.From.UserName,
		update.Message.Text,
		err.Error(),
	)

	errMsg := tgbotapi.NewMessage(config.GetConfig().MillionDebillion, text)
	_, err = GetBot().Send(errMsg)
	if err != nil {
		log.Fatal(err)
	}
}
