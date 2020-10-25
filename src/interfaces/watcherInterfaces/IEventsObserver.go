package watcherinterfaces

import "leagueapi.com.br/brain/src/broadcast"

// IObserverEvents interface to observe events
type IObserverEvents interface {
	Update(observerUpdate chan broadcast.Player, player *broadcast.Player)
	GetID() string
}
