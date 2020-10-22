package factoryevents

import (
	"leagueapi.com.br/brain/src/events"
	IWatcher "leagueapi.com.br/brain/src/interfaces/watcherInterfaces"
)

// CreateUserEventFactory factory to create my event of create user
type CreateUserEventFactory struct{}

// CreateEvent implement of interface. abstract factory
func (c *CreateUserEventFactory) CreateEvent() IWatcher.IEvents {
	return events.NewCreatePlayerEvent()
}

// NewCreateUserEventFactory constructor
func NewCreateUserEventFactory() *CreateUserEventFactory {
	return &CreateUserEventFactory{}
}
