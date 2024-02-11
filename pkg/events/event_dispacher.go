package events

import "errors"

var ErrorAlreadyRegistered = errors.New("Handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHanderInterface
}

func NewEventDispacher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHanderInterface),
	}
}

func (ed *EventDispatcher) Register(name string, handle EventHanderInterface) error {
	if _, ok := ed.handlers[name]; !ok {
		for _, h := range ed.handlers[name] {
			if h == handle {
				return ErrorAlreadyRegistered
			}
		}
	}

	ed.handlers[name] = append(ed.handlers[name], handle)
	return nil
}
