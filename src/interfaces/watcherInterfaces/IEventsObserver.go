package watcherinterfaces

// IObserverEvents interface to observe events
type IObserverEvents interface {
	Update(observerUpdate chan string)
	GetID() string
}
