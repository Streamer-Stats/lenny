package entities

import (
	"leagueapi.com.br/brain/src/interfaces"
	interfacefactory "leagueapi.com.br/brain/src/interfaces/factory"
)

// Events is a watcher
type Events struct {
	Events         []string
	EventsRegister []interfaces.IEvents
}

// RegisterEvents register new events
func (b *Events) RegisterEvents() *Events {
	var factory interfacefactory.IEventsFactory
	var err error
	for _, event := range b.Events {
		factory, err = interfacefactory.GetEventsFactory(event)
		if err == nil {
			b.EventsRegister = append(b.EventsRegister, factory.CreateEvent(event))
		}
	}

	return b
}

// NewEvents constructor
func NewEvents() *Events {
	return &Events{
		Events: []string{"registerUser"},
	}
}
