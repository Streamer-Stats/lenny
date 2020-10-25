package handlers

import (
	"leagueapi.com.br/brain/src/broadcast"
	"leagueapi.com.br/brain/src/events"
)

// PlayerHandler handle of createPlayerEvent
type PlayerHandler struct {
	event   *events.CreatePlayerEvent
	watcher *Watcher
}

// CreateNewPlayer create a new player watcher
func (handler *PlayerHandler) CreateNewPlayer(username string) {
	playerObserver := handler.watcher.PlayerWatcher(username, handler.event.ID)
	handler.event.Register(playerObserver)
	playerObserver.CheckIfPlayerExist()

}

// NotifyAll notify all watchers
func (handler *PlayerHandler) NotifyAll(observerUpdate chan broadcast.Player) {
	handler.event.NotifyAll(observerUpdate)
}

// NewPlayerHandler constructor
func NewPlayerHandler(_event *events.CreatePlayerEvent, _watcher *Watcher) *PlayerHandler {
	return &PlayerHandler{
		event:   _event,
		watcher: _watcher,
	}
}
