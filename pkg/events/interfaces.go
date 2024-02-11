package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateAndTime() time.Time
	GetPayload() interface{}
}

type EventHanderInterface interface {
	Handle(event EventInterface)
}

type EventDispatcherInterface interface {
	Register(name string, handle EventHanderInterface) error
	Dispach(event EventInterface) error
	Remove(name string, handle EventHanderInterface) error
	Has(name string) bool
	Clear() error
}
