package entities

import (
	"fmt"
)

// Brain is a watcher
type Brain struct {
	events *Events
}

// Handle handle the decisions
func (brain *Brain) Handle() {
	newPlayerEvent := brain.events.GetCreatePlayerEvent()
	fmt.Println(newPlayerEvent.ID)
}

// NewBrain constructor
func NewBrain() *Brain {
	return &Brain{
		events: NewEvents().RegisterEvents(),
	}
}
