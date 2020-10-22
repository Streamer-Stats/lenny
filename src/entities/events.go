package entities

import (
	"leagueapi.com.br/brain/src/enum"
	event "leagueapi.com.br/brain/src/events"
	abstractfactory "leagueapi.com.br/brain/src/interfaces/factory"
	IWatcher "leagueapi.com.br/brain/src/interfaces/watcherInterfaces"
)

// Events is a watcher
type Events struct {
	Events         []enum.EventsEnum
	eventsRegister map[enum.EventsEnum]IWatcher.IEvents
}

// RegisterEvents register new events
func (events *Events) RegisterEvents() *Events {
	var factory abstractfactory.IEventsFactory
	var err error
	for _, event := range events.Events {
		factory, err = abstractfactory.GetEventsFactory(event)
		if err == nil {
			events.eventsRegister[event] = factory.CreateEvent()
		}
	}

	return events
}

func (events *Events) get(event enum.EventsEnum) IWatcher.IEvents {

	switch event {
	case enum.CreateNewPlayer:
		return events.eventsRegister[event]
	default:
		return nil
	}

}

// GetCreatePlayerEvent get create player event
func (events *Events) GetCreatePlayerEvent() *event.CreatePlayerEvent {
	return events.get(enum.CreateNewPlayer).(*event.CreatePlayerEvent)
}

// NewEvents constructor
func NewEvents() *Events {
	return &Events{
		Events:         []enum.EventsEnum{enum.CreateNewPlayer},
		eventsRegister: make(map[enum.EventsEnum]IWatcher.IEvents),
	}
}
