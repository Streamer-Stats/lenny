package interfaces

// IObserverEvents interface to observe events
type IObserverEvents interface {
	Update(t string)
	GetID() string
}
