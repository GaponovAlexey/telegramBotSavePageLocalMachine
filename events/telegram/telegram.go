package telegram

import (
	"fmt"

	"tg/sitesess.ca/client/telegram"
	"tg/sitesess.ca/events"
	"tg/sitesess.ca/storage"

)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      client,
		storage: storage,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	update, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, fmt.Errorf("Fetch %s ", err)
	}
	res := make([]events.Event, 0, len(update))

	for _, u := range update {
		res = append(res, event(u))
	}
}

func event(upd telegram.Update) events.Event {
	res:= events.Event {
		Type: fetchType(upd),
		Text: fetchText(upd),
	}
}

func fetchText(upd telegram.Update) string {
	return upd.Message
}

func fetchType(upd telegram.Update) events.Type {

}