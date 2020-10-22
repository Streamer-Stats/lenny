package interfacefactory

import (
	"fmt"

	"leagueapi.com.br/brain/src/enum"
	factoryevents "leagueapi.com.br/brain/src/factory/events"
	"leagueapi.com.br/brain/src/interfaces"
)

// IObserverFactory interface of factory
type IObserverFactory interface {
	CreateObserver(id string) interfaces.IObserverEvents
}

// GetObserverFactory static function to construct the IObserverFactory
func GetObserverFactory(event enum.EventsEnum) IObserverFactory {
	switch event {
	case enum.CreateNewPlayer:
		return factoryevents.NewCreateUserEventFactory(), nil
	default:
		return nil, fmt.Errorf("Wrong event type passed")
	}
}
