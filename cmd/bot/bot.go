package bot

import (
	"log"

	"ria-bot/configs"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(configs.TELEGRAM_API_KEY)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return bot
}
