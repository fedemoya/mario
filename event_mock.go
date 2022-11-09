package mario

import (
	"github.com/stretchr/testify/mock"
)

type EventMockVisitor interface{}

type EventMock struct {
	mock.Mock
}

func (e *EventMock) Accept(visitor EventMockVisitor) error {
	args := e.Called(visitor)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (e *EventMock) Ack() error {
	args := e.Called()
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (e EventMock) Nack(opts interface{}) error {
	args := e.Called(opts)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
