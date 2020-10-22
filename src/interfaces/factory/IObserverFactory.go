package interfacefactory

import (
	"fmt"

	observersfactory "leagueapi.com.br/brain/src/factory/observers"

	"leagueapi.com.br/brain/src/enum"
	IWatcher "leagueapi.com.br/brain/src/interfaces/watcherInterfaces"
)

// IObserverFactory interface of factory
type IObserverFactory interface {
	CreateObserver(id string) IWatcher.IObserverEvents
}

// GetObserverFactory static function to construct the IObserverFactory
func GetObserverFactory(event enum.EventsEnum) (IObserverFactory, error) {
	switch event {
	case enum.CreateNewPlayer:
		return observersfactory.NewObserverFactory(), nil
	default:
		return nil, fmt.Errorf("Wrong event type passed")
	}
}
