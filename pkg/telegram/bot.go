package telegram

import (
	"fmt"
	"log"

	"ria-bot/configs"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI

	//messages config.Messages
}

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

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot: bot,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// Handle commands
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				fmt.Println(err)
			}

			continue
		}

		// Handle regular messages
		if err := b.handleMessage(update.Message); err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
