package observersfactory

import (
	IWatcher "leagueapi.com.br/brain/src/interfaces/watcherInterfaces"
	"leagueapi.com.br/brain/src/observers"
)

// ObserverFactory factory of observers
type ObserverFactory struct{}

// CreateObserver create a observer
func (o *ObserverFactory) CreateObserver(id string) IWatcher.IObserverEvents {
	return observers.NewPlayerObserver(id)
}

// NewObserverFactory constructor
func NewObserverFactory() *ObserverFactory {
	return &ObserverFactory{}
}
