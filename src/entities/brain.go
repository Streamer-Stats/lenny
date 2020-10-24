package entities

import (
	"strings"

	"leagueapi.com.br/brain/src/infrastructure/syncronize"

	"leagueapi.com.br/brain/src/handlers"
)

// Brain is a watcher
type Brain struct {
	events  *Events
	handler *handlers.Handlers
	sync    *syncronize.Syncronize
}

func (brain *Brain) AddSyncronize(syncronize *syncronize.Syncronize) *Brain {
	brain.sync = syncronize
	return brain
}

// StartHandlers initiate all handlers
func (brain *Brain) StartHandlers() *Brain {
	brain.handler.PlayerHandler(brain.events.GetCreatePlayerEvent())
	return brain
}

// Handle handle the decisions
func (brain *Brain) Handle(command string) {
	switch strings.TrimSpace(command) {
	case "CREATEUSER":
		go brain.handler.NewPlayerHandler.NotifyAll(brain.sync.BroadCast)
	default:
		go brain.handler.NewPlayerHandler.CreateNewPlayer(command)
	}

}

// NewBrain constructor
func NewBrain() *Brain {
	return &Brain{
		events:  NewEvents().RegisterEvents(),
		handler: handlers.NewHandler(),
	}
}
