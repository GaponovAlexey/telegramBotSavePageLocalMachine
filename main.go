package main

import (
	"flag"
	"log"

	tgClient "tg/sitesess.ca/client/telegram"
	"tg/sitesess.ca/consumer/e-consumers"
	"tg/sitesess.ca/events/telegram"
	"tg/sitesess.ca/storage/files"

)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files-storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Println("start")

	consumer := e_consumers.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("Server is stopped", err)
	}

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"5600797009:AAHNMQbKM5v0rvm1SNFCsSh_xh72ZZ8-Agg",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("Token is not specified")
	}

	return *token
}
