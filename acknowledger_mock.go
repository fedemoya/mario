package mario

import "github.com/stretchr/testify/mock"

type AcknowledgerMock struct {
	mock.Mock
}

func (a *AcknowledgerMock) Ack() error {
	args := a.Called()
	return args.Error(0)
}

func (a *AcknowledgerMock) Nack(opts interface{}) error {
	args := a.Called(opts)
	return args.Error(0)
}
