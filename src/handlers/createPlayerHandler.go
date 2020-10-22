package handlers

import (
	"leagueapi.com.br/brain/src/events"
)

// CreatePlayerHandler handle of createPlayerEvent
type CreatePlayerHandler struct {
	event   *events.CreatePlayerEvent
	watcher *Watcher
}

// CreateNewPlayer create a new player watcher
func (handler *CreatePlayerHandler) CreateNewPlayer(username string) {
	handler.event.Register(handler.watcher.CreateNewPlayerWatcher(username, handler.event.ID))
}

// NotifyAll notify all watchers
func (handler *CreatePlayerHandler) NotifyAll() {
	handler.event.NotifyAll()
}

// NewCreatePlayerHandler constructor
func NewCreatePlayerHandler(_event *events.CreatePlayerEvent, _watcher *Watcher) *CreatePlayerHandler {
	return &CreatePlayerHandler{
		event:   _event,
		watcher: _watcher,
	}
}
