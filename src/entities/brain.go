package entities

import (
	"leagueapi.com.br/brain/src/interfaces"
	interfacefactory "leagueapi.com.br/brain/src/interfaces/factory"
)

// Brain is a watcher
type Brain struct {
	events         []string
	eventsRegister []interfaces.IEvents
}

// RegisterEvents register new events
func (b *Brain) RegisterEvents() *Brain {
	var factory interfacefactory.IEventsFactory
	var err error
	for _, event := range b.events {
		factory, err = interfacefactory.GetEventsFactory(event)
		if err == nil {
			b.eventsRegister = append(b.eventsRegister, factory.CreateEvent())
		}
	}

	return b
}

// NewBrain constructor
func NewBrain() *Brain {
	return &Brain{
		events: []string{"registerUser"},
	}
}
