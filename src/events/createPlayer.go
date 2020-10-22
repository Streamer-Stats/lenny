package events

import (
	"leagueapi.com.br/brain/src/enum"
	IWatcher "leagueapi.com.br/brain/src/interfaces/watcherInterfaces"
	"leagueapi.com.br/brain/src/observers"
)

// CreatePlayerEvent event of player criation
type CreatePlayerEvent struct {
	ID           enum.EventsEnum
	ObserverList []*observers.CreatePlayerObserver
}

// Register register a new type event player
func (e *CreatePlayerEvent) Register(observer IWatcher.IObserverEvents) {
	e.ObserverList = append(e.ObserverList, observer.(*observers.CreatePlayerObserver))
}

// Deregister deregister a new type event player
func (e *CreatePlayerEvent) Deregister(observer IWatcher.IObserverEvents) {}

// NotifyAll notify all watcher
func (e *CreatePlayerEvent) NotifyAll() {
	for _, observers := range e.ObserverList {
		observers.Update()
	}
}

// ReturnID return eventID
func (e *CreatePlayerEvent) ReturnID() enum.EventsEnum {
	return e.ID
}

// NewCreatePlayerEvent constructor
func NewCreatePlayerEvent() *CreatePlayerEvent {
	return &CreatePlayerEvent{
		ID: enum.CreateNewPlayer,
	}
}
