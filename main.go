package main

import (
	"flag"
	"log"
	"tg/sitesess-ca/consumer/e-consumers"
	"tg/sitesess-ca/events/telegram"
	"tg/sitesess-ca/storage/files"
	tgClient "tg/sitesess-ca/client/telegram"

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
		"", // перед компиляцией нужно закинуть токен 
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("Token is not specified")
	}

	return *token
}
