package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name        string
	DateAndTime time.Time
	Payload     interface{}
}

type TestEventHandler struct {
	ID int
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handle2         TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetDateAndTime() time.Time {
	return e.DateAndTime
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (h *TestEventHandler) Handle(event EventInterface) {
	// do something
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.event = TestEvent{
		Name:        "test",
		DateAndTime: time.Now(),
		Payload:     "test",
	}

	suite.event2 = TestEvent{
		Name:        "test2",
		DateAndTime: time.Now(),
		Payload:     "test2",
	}

	suite.handler = TestEventHandler{ID: 1}
	suite.handle2 = TestEventHandler{ID: 2}
	suite.handler3 = TestEventHandler{ID: 3}
	suite.eventDispatcher = NewEventDispacher()

}

func (suite *EventDispatcherTestSuite) TestEventHandler() {
	assert.True(suite.T(), true)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
