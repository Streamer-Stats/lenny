package entities

import (
	"leagueapi.com.br/brain/src/enum"
	abstractfactory "leagueapi.com.br/brain/src/interfaces/factory"
	IWatcher "leagueapi.com.br/brain/src/interfaces/watcherInterfaces"
)

// Events is a watcher
type Events struct {
	Events         []enum.EventsEnum
	EventsRegister []IWatcher.IEvents
}

// RegisterEvents register new events
func (e *Events) RegisterEvents() *Events {
	var factory abstractfactory.IEventsFactory
	var err error
	for _, event := range e.Events {
		factory, err = abstractfactory.GetEventsFactory(event)
		if err == nil {
			e.EventsRegister = append(e.EventsRegister, factory.CreateEvent())
		}
	}

	return e
}

// NewEvents constructor
func NewEvents() *Events {
	return &Events{
		Events: []enum.EventsEnum{enum.CreateNewPlayer},
	}
}
