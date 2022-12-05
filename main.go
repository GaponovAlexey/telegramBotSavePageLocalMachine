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
	storagePath = "storage"
)

func main() {

	eventsProcesor := telegram.New(tgClient.New(tgBotHost, mustToken()), files.New(storagePath))

	log.Println("start")

	consumer := e_consumers.New()

	// processor = processor.New(thClient)

	// consumer.Start(fetcher, processor)

}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"111222",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("Token is not specified")
	}

	return *token
}
