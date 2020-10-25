package handlers

import (
	"leagueapi.com.br/brain/src/events"
)

type Handlers struct {
	watcher          *Watcher
	NewPlayerHandler *PlayerHandler
}

func (handlers *Handlers) PlayerHandler(event *events.CreatePlayerEvent) {
	handlers.NewPlayerHandler = NewPlayerHandler(event, handlers.watcher)
}

func NewHandler() *Handlers {
	return &Handlers{
		watcher: NewWatcher(),
	}
}
