package memdb

import (
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"mario"
	"testing"
	"time"
)

type RepositoryTestSuite struct {
	suite.Suite

	db *memdb.MemDB

	cloudEventID            string
	cloudEventSource        string
	cloudEventType          string
	cloudEventTime          int64
	cloudEventCorrelationID string
	cloudEventData          []byte

	repository mario.CloudEventRepository
	cloudEvent mario.CloudEvent
	err        error
}

func (s *RepositoryTestSuite) TestRepository_Add() {

	s.given().
		aDB().and().
		aRepository().and().
		aCloudEvent()

	s.when().
		theCloudEventIsAdded()

	s.then().
		thereIsNoError().and().
		theSavedCloudEventHasCorrectValues()
}

func TestRepository_Stream(t *testing.T) {

	db := InitDB()

	id := uuid.New()
	source := "paymentapi"
	eventType := "withdrawal.created"
	time := time.Now().Unix()
	correlationID := uuid.New()
	data := []byte("some nice json object")

	mock := &mario.CloudEventMock{}
	mock.On("ID").Return(id.String())
	mock.On("Source").Return(source)
	mock.On("Type").Return(eventType)
	mock.On("Time").Return(time)
	mock.On("CorrelationID").Return(correlationID.String())
	mock.On("Data").Return(data, nil)

	repository := NewRepository(db, mario.NewCloudEventBuilderImpl())
	err := repository.Add(mock)
	require.NoError(t, err)

	ch, err := repository.Stream("")
	require.NoError(t, err)

	cloudEvent := <-ch

	require.NotNil(t, cloudEvent)
	require.Equal(t, id.String(), cloudEvent.ID())
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *RepositoryTestSuite) given() *RepositoryTestSuite {
	return s
}

func (s *RepositoryTestSuite) when() *RepositoryTestSuite {
	return s
}

func (s *RepositoryTestSuite) then() *RepositoryTestSuite {
	return s
}

func (s *RepositoryTestSuite) and() *RepositoryTestSuite {
	return s
}

func (s *RepositoryTestSuite) aCloudEvent() *RepositoryTestSuite {

	id := uuid.New()
	source := "paymentapi"
	eventType := "withdrawal.created"
	time := time.Now().Unix()
	correlationID := uuid.New()
	data := []byte("some nice json object")

	mock := &mario.CloudEventMock{}
	mock.On("ID").Return(id.String())
	mock.On("Source").Return(source)
	mock.On("Type").Return(eventType)
	mock.On("Time").Return(time)
	mock.On("CorrelationID").Return(correlationID.String())
	mock.On("Data").Return(data, nil)

	s.cloudEvent = mock

	return s
}

func (s *RepositoryTestSuite) thereIsNoError() *RepositoryTestSuite {
	s.Require().NoError(s.err)
	return s
}

func (s *RepositoryTestSuite) aRepository() *RepositoryTestSuite {
	s.repository = NewRepository(s.db, mario.NewCloudEventBuilderImpl())
	return s
}

func (s *RepositoryTestSuite) theCloudEventIsAdded() *RepositoryTestSuite {
	s.err = s.repository.Add(s.cloudEvent)
	return s
}

func (s *RepositoryTestSuite) theSavedCloudEventHasCorrectValues() {
	txn := s.db.Txn(false)
	resultIter, err := txn.Get("events", "source", s.cloudEventSource)
	defer txn.Abort()

	s.Require().NoError(err)

	row := resultIter.Next()
	s.Require().NotNil(row)

	storableEvent, ok := row.(StorableCloudEvent)
	s.Require().True(ok)

	s.Require().Equal(storableEvent.ID, s.cloudEventID)
	s.Require().Equal(storableEvent.Data, s.cloudEventData)
}

func (s *RepositoryTestSuite) aDB() *RepositoryTestSuite {
	s.db = InitDB()
	return s
}
