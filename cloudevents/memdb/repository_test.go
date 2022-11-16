package memdb

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRepository_Add(t *testing.T) {

	id := uuid.New()
	source := "paymentapi"
	eventType := "withdrawal.created"
	time := time.Now().Unix()
	correlationID := uuid.New()
	data := []byte("some nice json object")

	mock := &SerializableCloudEventMock{}
	mock.On("ID").Return(id.String())
	mock.On("Source").Return(source)
	mock.On("Type").Return(eventType)
	mock.On("Time").Return(time)
	mock.On("CorrelationID").Return(correlationID.String())
	mock.On("Serialize").Return(data, nil)

	memdbRepo := NewRepository()
	err := memdbRepo.Add(mock)

	require.NoError(t, err)
}

type SerializableCloudEventMock struct {
	mock.Mock
}

func (mock *SerializableCloudEventMock) ID() string {
	args := mock.Called()
	return args.String(0)
}

func (mock *SerializableCloudEventMock) Source() string {
	args := mock.Called()
	return args.String(0)
}

func (mock *SerializableCloudEventMock) Type() string {
	args := mock.Called()
	return args.String(0)
}

func (mock *SerializableCloudEventMock) Time() int64 {
	args := mock.Called()
	return args.Get(0).(int64)
}

func (mock *SerializableCloudEventMock) CorrelationID() string {
	args := mock.Called()
	return args.String(0)
}

func (mock *SerializableCloudEventMock) Serialize() ([]byte, error) {
	args := mock.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), nil
}
