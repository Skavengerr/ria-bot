package telegram

import (
	"fmt"
	"log"

	configs "ria-bot/configs"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	cfg *configs.Config
}

func InitBot(configs *configs.Config) *tgbotapi.BotAPI {
	fmt.Println("asdads", configs)
	bot, err := tgbotapi.NewBotAPI(configs.TelegramApiKey)
	if err != nil {
		log.Panic("dasdas", err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return bot
}

func NewBot(bot *tgbotapi.BotAPI, cfg *configs.Config) *Bot {
	return &Bot{
		bot: bot,
		cfg: cfg,
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
