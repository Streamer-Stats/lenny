package watcherinterfaces

// IObserverEvents interface to observe events
type IObserverEvents interface {
	Update()
	GetID() string
}
