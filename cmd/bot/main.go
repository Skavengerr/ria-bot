package main

import (
	"log"
	"ria-bot/cmd/bot"
	"ria-bot/pkg/telegram"
)

func main() {
	botApi := bot.InitBot()

	bot := telegram.NewBot(botApi)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
