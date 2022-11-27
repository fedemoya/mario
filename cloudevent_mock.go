package mario

import "github.com/stretchr/testify/mock"

type CloudEventMock struct {
	mock.Mock
}

func (mock CloudEventMock) ID() string {
	args := mock.Called()
	return args.String(0)
}

func (mock CloudEventMock) Source() string {
	args := mock.Called()
	return args.String(0)
}

func (mock CloudEventMock) Type() string {
	args := mock.Called()
	return args.String(0)
}

func (mock CloudEventMock) Time() int64 {
	args := mock.Called()
	return args.Get(0).(int64)
}

func (mock CloudEventMock) CorrelationID() string {
	args := mock.Called()
	return args.String(0)
}

func (mock CloudEventMock) Status() CloudEventStatus {
	args := mock.Called()
	return args.Get(0).(CloudEventStatus)
}

func (mock CloudEventMock) Data() []byte {
	args := mock.Called()
	return args.Get(0).([]byte)
}
