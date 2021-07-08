package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updatesChannel, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updatesChannel {
		if update.Message == nil {
			continue
		}

		log.Printf("Received message from [%s]: [%s]",
			update.Message.From.UserName, update.Message.Text)

		reply := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		reply.ReplyToMessageID = update.Message.MessageID

		_, err := bot.Send(reply)
		if err != nil {
			log.Println("Error while sending message: ", err)
		}
	}
}
