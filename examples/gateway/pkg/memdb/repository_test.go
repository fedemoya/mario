package memdb

import (
	"encoding/json"
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
	mock.On("Data").Return(data, nil)

	db := InitDB()

	memdbRepo := NewRepository(db)
	err := memdbRepo.Add(mock)

	require.NoError(t, err)

	txn := db.Txn(false)
	resultIter, err := txn.Get("events", "source", source)
	defer txn.Abort()

	require.NoError(t, err)

	row := resultIter.Next()
	require.NotNil(t, row)

	storableEvent, ok := row.(StorableCloudEvent)
	require.True(t, ok)

	require.Equal(t, storableEvent.IDField, id.String())
	require.Equal(t, storableEvent.DataField, json.RawMessage(data))
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

func (mock *SerializableCloudEventMock) Data() []byte {
	args := mock.Called()
	return args.Get(0).([]byte)
}
