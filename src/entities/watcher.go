package entities

import (
	"leagueapi.com.br/brain/src/enum"
	abstractfactory "leagueapi.com.br/brain/src/interfaces/factory"
)

// Watcher is a facade of watcher
type Watcher struct {
}

// CreateNewWatcher create a new watcher
func (w *Watcher) CreateNewWatcher(event enum.EventsEnum) abstractfactory.IObserverFactory {
	factory, err := abstractfactory.GetObserverFactory(event)
	if err != nil {
		panic(err)
	}
	return factory
}

// NewWatcher constructor
func NewWatcher() *Watcher {
	return &Watcher{}
}
