package events

import (
	"errors"
	"sync"
)

var ErrorAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHanderInterface
}

func NewEventDispacher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHanderInterface),
	}
}

func (ed *EventDispatcher) Register(name string, handle EventHanderInterface) error {
	if _, ok := ed.handlers[name]; ok {
		for _, h := range ed.handlers[name] {
			if h == handle {
				return ErrorAlreadyRegistered
			}
		}
	}

	ed.handlers[name] = append(ed.handlers[name], handle)
	return nil
}

func (ed *EventDispatcher) Clear() error {
	ed.handlers = make(map[string][]EventHanderInterface)
	return nil
}

func (ed *EventDispatcher) Has(name string, handle EventHanderInterface) bool {
	if _, ok := ed.handlers[name]; ok {
		for _, h := range ed.handlers[name] {
			if h == handle {
				return true
			}
		}
	}

	return false
}

func (ed *EventDispatcher) Remove(name string, handle EventHanderInterface) error {
	if _, ok := ed.handlers[name]; ok {
		for i, h := range ed.handlers[name] {
			if h == handle {
				ed.handlers[name] = append(ed.handlers[name][:i], ed.handlers[name][i+1:]...)
				return nil
			}
		}
	}

	return nil
}

func (ed *EventDispatcher) Dispach(event EventInterface) error {
	var wg sync.WaitGroup
	if _, ok := ed.handlers[event.GetName()]; ok {
		wg.Add(len(ed.handlers))
		for _, h := range ed.handlers[event.GetName()] {
			go h.Handle(event, &wg)
		}
		wg.Wait()
	}

	return nil
}
