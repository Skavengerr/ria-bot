package telegram

import (
	"log"

	"ria-bot/api/server"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart       = "start"
	start              = "Hi, choose what you need to find"
	startMake          = "Hi, type a car model that you want to find. Ex: Audi"
	unknownCommand     = "I don't know this command"
	keyboardShowAll    = "showAll"
	keyboardShowByMark = "showByMark"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/"+keyboardShowAll),
		tgbotapi.NewKeyboardButton("/"+keyboardShowByMark),
	),
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case keyboardShowAll:
		return b.handleGetAllCars(message)
	case keyboardShowByMark:
		return b.handleGetCarsByMark(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, start)

	msg.ReplyMarkup = numericKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	cars := server.GetCarsByMark(message.Text)

	for _, car := range cars {
		msg := tgbotapi.NewMessage(message.Chat.ID, car)

		_, err := b.bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (b *Bot) handleGetAllCars(message *tgbotapi.Message) error {
	cars := server.GetAllCars()

	for _, car := range cars {
		msg := tgbotapi.NewMessage(message.Chat.ID, car)

		_, err := b.bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (b *Bot) handleGetCarsByMark(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Type mark that you want to find")
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, unknownCommand)
	_, err := b.bot.Send(msg)
	return err
}
