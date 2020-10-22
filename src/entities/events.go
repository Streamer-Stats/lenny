package entities

import (
	"leagueapi.com.br/brain/src/interfaces"
	interfacefactory "leagueapi.com.br/brain/src/interfaces/factory"
)

// Events is a watcher
type Events struct {
	events         []string
	eventsRegister []interfaces.IEvents
}

// RegisterEvents register new events
func (b *Events) RegisterEvents() *Events {
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

// NewEvents constructor
func NewEvents() *Events {
	return &Events{
		events: []string{"registerUser"},
	}
}
