package watcherinterfaces

import (
	"leagueapi.com.br/brain/src/enum"
)

// IEvents events objs
type IEvents interface {
	Register(observer IObserverEvents)
	Deregister(observer IObserverEvents)
	ReturnID() enum.EventsEnum
	NotifyAll()
}
