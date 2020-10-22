package observersfactory

import (
	"leagueapi.com.br/brain/src/interfaces"
	"leagueapi.com.br/brain/src/observers"
)

// ObserverFactory factory of observers
type ObserverFactory struct{}

// CreateObserver create a observer
func (o *ObserverFactory) CreateObserver(id string) interfaces.IObserverEvents {
	return observers.NewCreatePlayerObserver(id)
}

// NewObserverFactory constructor
func NewObserverFactory() *ObserverFactory {
	return &ObserverFactory{}
}
