package telegram

import "tg/sitesess.ca/client/telegram"

type Processor struct {
	tg *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client) {
	
}