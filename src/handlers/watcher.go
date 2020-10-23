package handlers

import (
	"leagueapi.com.br/brain/src/enum"
	abstractfactory "leagueapi.com.br/brain/src/interfaces/factory"
	"leagueapi.com.br/brain/src/observers"
)

// Watcher is a facade of watcher
type Watcher struct {
}

// CreateNewWatcher create a new watcher
func (watcher *Watcher) createNewWatcher(event enum.EventsEnum) abstractfactory.IObserverFactory {
	factory, err := abstractfactory.GetObserverFactory(event)
	if err != nil {
		panic(err)
	}
	return factory
}

// CreateNewPlayerWatcher create a new player watcher
func (watcher *Watcher) CreateNewPlayerWatcher(username string, event enum.EventsEnum) *observers.CreatePlayerObserver {
	return watcher.createNewWatcher(event).CreateObserver(username).(*observers.CreatePlayerObserver)
}

// NewWatcher constructor
func NewWatcher() *Watcher {
	return &Watcher{}
}
