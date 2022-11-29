package memdb

import (
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
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
	cloudEventStatus        mario.CloudEventStatus
	cloudEventData          []byte

	repository     mario.CloudEventRepository
	cloudEventMock *mario.CloudEventMock
	err            error

	savedCloudEvent mario.CloudEvent
}

func (s *RepositoryTestSuite) BeforeTest(_, _ string) {
	s.db = InitDB()
	println("BeforeTest")
}

func (s *RepositoryTestSuite) TestRepository_Add() {
	s.given().
		aRepository().and().
		aCloudEvent()

	s.when().
		theCloudEventIsAddedToTheRepository()

	s.then().
		thereIsNoError().and().
		theSavedStorableCloudEventHasExpectedValues()
}

func (s *RepositoryTestSuite) TestRepository_Stream() {
	s.given().
		aRepository().and().
		aCloudEvent().and().
		theCloudEventIsAddedToTheRepository()

	s.when().
		theCloudEventIsConsumedFromTheRepositoryStream()

	s.then().
		thereIsNoError().and().
		theSavedCloudEventHasCorrectValues()
}

func (s *RepositoryTestSuite) TestRepository_UpdateStatus() {
	s.given().
		aRepository().and().
		aCloudEvent().and().
		theCloudEventIsAddedToTheRepository()

	s.when().
		theCloudEventStatusIsUpdatedTo(mario.CloudEventProcessed)

	s.then().
		thereIsNoError().and().
		theSavedStorableCloudEventHasExpectedValues()
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

	s.cloudEventID = uuid.New().String()
	s.cloudEventSource = "paymentapi"
	s.cloudEventType = "withdrawal.created"
	s.cloudEventTime = time.Now().Unix()
	s.cloudEventCorrelationID = uuid.New().String()
	s.cloudEventStatus = mario.CloudEventPending
	s.cloudEventData = []byte("some nice json object")

	mock := &mario.CloudEventMock{}
	mock.On("ID").Return(s.cloudEventID)
	mock.On("Source").Return(s.cloudEventSource)
	mock.On("Type").Return(s.cloudEventType)
	mock.On("Time").Return(s.cloudEventTime)
	mock.On("CorrelationID").Return(s.cloudEventCorrelationID)
	mock.On("Status").Return(s.cloudEventStatus)
	mock.On("Data").Return(s.cloudEventData, nil)

	s.cloudEventMock = mock

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

func (s *RepositoryTestSuite) theCloudEventIsAddedToTheRepository() *RepositoryTestSuite {
	s.err = s.repository.Add(s.cloudEventMock)
	return s
}

func (s *RepositoryTestSuite) theSavedStorableCloudEventHasExpectedValues() {
	txn := s.db.Txn(false)
	resultIter, err := txn.Get("events", "id", s.cloudEventID)
	defer txn.Abort()

	s.Require().NoError(err)

	row := resultIter.Next()
	s.Require().NotNil(row)

	storableEvent, ok := row.(StorableCloudEvent)
	s.Require().True(ok)

	s.Require().Equal(storableEvent.ID, s.cloudEventID)
	s.Require().Equal(storableEvent.Source, s.cloudEventSource)
	s.Require().Equal(storableEvent.Type, s.cloudEventType)
	s.Require().Equal(storableEvent.Time, s.cloudEventTime)
	s.Require().Equal(storableEvent.CorrelationID, s.cloudEventCorrelationID)
	s.Require().Equal(storableEvent.Status, s.cloudEventStatus)
	s.Require().Equal(storableEvent.Data, s.cloudEventData)
}

func (s *RepositoryTestSuite) theCloudEventIsConsumedFromTheRepositoryStream() *RepositoryTestSuite {
	ch, err := s.repository.Stream("")
	s.Require().NoError(err)

	s.savedCloudEvent = <-ch

	return s
}

func (s *RepositoryTestSuite) theSavedCloudEventHasCorrectValues() *RepositoryTestSuite {

	s.Require().Equal(s.cloudEventID, s.savedCloudEvent.ID())
	s.Require().Equal(s.cloudEventData, s.savedCloudEvent.Data())

	return s
}

func (s *RepositoryTestSuite) theCloudEventStatusIsUpdatedTo(processed mario.CloudEventStatus) {
	s.cloudEventStatus = processed
	s.err = s.repository.UpdateStatus(s.cloudEventMock, processed)
}
