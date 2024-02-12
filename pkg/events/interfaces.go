package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateAndTime() time.Time
	GetPayload() interface{}
}

type EventHanderInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(name string, handle EventHanderInterface) error
	Dispach(event EventInterface) error
	Remove(name string, handle EventHanderInterface) error
	Has(name string, handle EventHanderInterface) bool
	Clear() error
}
