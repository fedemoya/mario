package mario

import "github.com/stretchr/testify/mock"

type AcknowledgementHandlerMock[EventVisitor any] struct {
	mock.Mock
}

func (a *AcknowledgementHandlerMock[V]) OnSuccess(acknowledger Acknowledger) error {
	args := a.Called(acknowledger)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (a *AcknowledgementHandlerMock[V]) OnError(acknowledger Acknowledger, err error) error {
	args := a.Called(acknowledger, err)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
