package entities

import (
	"fmt"

	"leagueapi.com.br/brain/src/enum"
	"leagueapi.com.br/brain/src/events"
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

// Get get specific event
func (e *Events) Get(event enum.EventsEnum) (interface{}, error) {
	var eventType interface{}
	var err error

	for _, register := range e.EventsRegister {
		if event == register.ReturnID() {
			switch event {
			case enum.CreateNewPlayer:
				eventType = register.(*events.CreatePlayerEvent)
				err = nil
				break
			default:
				eventType = nil
				err = fmt.Errorf("Have not register this event")
			}
		}
	}

	return eventType, err

}

// NewEvents constructor
func NewEvents() *Events {
	return &Events{
		Events: []enum.EventsEnum{enum.CreateNewPlayer},
	}
}
