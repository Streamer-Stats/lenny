package entities

import (
	"fmt"

	"leagueapi.com.br/brain/src/enum"
)

// Brain is a watcher
type Brain struct {
	events *Events
}

// Handle handle the decisions
func (brain *Brain) Handle() {
	newPlayerEvent, _ := brain.events.Get(enum.CreateNewPlayer)
	fmt.Println(newPlayerEvent)
}

// NewBrain constructor
func NewBrain() *Brain {
	return &Brain{
		events: NewEvents().RegisterEvents(),
	}
}
