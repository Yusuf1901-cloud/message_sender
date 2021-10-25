package repository

import (
	"github.com/Yusuf1901-cloud/message_sender/models"
	"github.com/Yusuf1901-cloud/message_sender/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var BotSenderRepo service.SenderI

func init() {
	bot, err := tgbotapi.NewBotAPI("2080627127:AAFaJuqZctfkBofFxdJNMaE32UK51YVHbJ0")
	if err != nil {
		panic(err)
	}

	BotSenderRepo = &botSenderRepo{
		bot: bot,
	}
}

type botSenderRepo struct {
	bot *tgbotapi.BotAPI
}

func (br *botSenderRepo) Send(msg *models.Message) error {
	tgMsg := tgbotapi.NewMessage(msg.ChatID, msg.Text)

	_, err := br.bot.Send(tgMsg)
	if err != nil {
		return err
	}

	return nil
}
