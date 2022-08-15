package main

import (
	"log"
	"ria-bot/pkg/telegram"

	configs "ria-bot/configs"
)

func main() {
	cfg, err := configs.InitViper(".")
	if err != nil {
		log.Fatal('a', err)
	}
	botApi := telegram.InitBot(&cfg)

	bot := telegram.NewBot(botApi, &cfg)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
