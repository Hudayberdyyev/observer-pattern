package interfaces

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}

type Observer interface {
	Update(string)
	GetId() string
}
