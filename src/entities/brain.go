package entities

import (
	"time"

	"leagueapi.com.br/brain/src/handlers"
)

// Brain is a watcher
type Brain struct {
	events  *Events
	handler *handlers.Handlers
}

// StartHandlers initiate all handlers
func (brain *Brain) StartHandlers() *Brain {
	brain.handler.RegisterCreatePlayerHandler(brain.events.GetCreatePlayerEvent())
	return brain
}

// Handle handle the decisions
func (brain *Brain) Handle() {
	brain.handler.NewPlayerHandler.CreateNewPlayer("BANOFFE")
	brain.handler.NewPlayerHandler.CreateNewPlayer("Frango SafaJhin")
	brain.handler.NewPlayerHandler.CreateNewPlayer("Zoiao")
	time.Sleep(5 * time.Second)
	brain.handler.NewPlayerHandler.NotifyAll()
}

// NewBrain constructor
func NewBrain() *Brain {
	return &Brain{
		events:  NewEvents().RegisterEvents(),
		handler: handlers.NewHandler(),
	}
}
