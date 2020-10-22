package handlers

import "leagueapi.com.br/brain/src/events"

type Handlers struct {
	watcher          *Watcher
	NewPlayerHandler *CreatePlayerHandler
}

func (handlers *Handlers) RegisterCreatePlayerHandler(event *events.CreatePlayerEvent) {
	handlers.NewPlayerHandler = NewCreatePlayerHandler(event, handlers.watcher)
}

func NewHandler() *Handlers {
	return &Handlers{
		watcher: NewWatcher(),
	}
}
