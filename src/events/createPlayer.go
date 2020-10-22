package events

import (
	"leagueapi.com.br/brain/src/interfaces"
	"leagueapi.com.br/brain/src/observers"
)

// CreatePlayerEvent event of player criation
type CreatePlayerEvent struct {
	ObserverList []*observers.CreatePlayerObserver
}

// Register register a new type event player
func (e *CreatePlayerEvent) Register(observer interfaces.IObserverEvents) {
	e.ObserverList = append(e.ObserverList, observer.(*observers.CreatePlayerObserver))
}

// Deregister deregister a new type event player
func (e *CreatePlayerEvent) Deregister(observer interfaces.IObserverEvents) {}

// NotifyAll notify all watcher
func (e *CreatePlayerEvent) NotifyAll() {
	for _, observers := range e.ObserverList {
		observers.Update("just for test")
	}
}

// NewCreatePlayerEvent constructor
func NewCreatePlayerEvent() *CreatePlayerEvent {
	return &CreatePlayerEvent{}
}
