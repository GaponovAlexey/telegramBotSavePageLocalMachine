package main

import (
	"flag"
	"log"

	"tg/sitesess.ca/client/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient = telegram.New(tgBotHost, mustToken())

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(thClient)

	// consumer.Start(fetcher, processor)

}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal(" token is not specified")
	}

	return *token
}
