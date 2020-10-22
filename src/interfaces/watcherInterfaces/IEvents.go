package watcherinterfaces

// IEvents events objs
type IEvents interface {
	Register(observer IObserverEvents)
	Deregister(observer IObserverEvents)
	NotifyAll()
}
