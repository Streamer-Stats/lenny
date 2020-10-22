package factoryevents

import (
	"leagueapi.com.br/brain/src/events"
	"leagueapi.com.br/brain/src/interfaces"
)

// CreateUserEventFactory factory to create my event of create user
type CreateUserEventFactory struct{}

// CreateEvent implement of interface. abstract factory
func (c *CreateUserEventFactory) CreateEvent(id string) interfaces.IEvents {
	return events.NewCreatePlayerEvent(id)
}

// NewCreateUserEventFactory constructor
func NewCreateUserEventFactory() *CreateUserEventFactory {
	return &CreateUserEventFactory{}
}
