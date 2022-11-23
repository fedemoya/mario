package memdb

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"mario"
	"time"
)

const (
	StorableEventStatusNotProcessed     = "not_processed"
	StorableEventStatusProcessed        = "processed"
	StorableEventStatusProcessingFailed = "processing_failed"
)

type Repository struct {
	db                *memdb.MemDB
	cloudEventBuilder mario.CloudEventBuilder
}

var _ mario.CloudEventRepository = (*Repository)(nil)

func NewRepository(db *memdb.MemDB, cloudEventBuilder mario.CloudEventBuilder) *Repository {
	return &Repository{db: db, cloudEventBuilder: cloudEventBuilder}
}

func (r *Repository) Add(event mario.CloudEvent) error {
	storableCloudEvent := StorableCloudEvent{
		ID:            event.ID(),
		Source:        event.Source(),
		CorrelationID: "",
		Type:          event.Type(),
		Time:          event.Time(),
		Data:          event.Data(),
		Status:        StorableEventStatusNotProcessed,
	}
	txn := r.db.Txn(true)
	err := txn.Insert("events", storableCloudEvent)
	if err != nil {
		txn.Abort()
		return fmt.Errorf("failed adding event with source %s and type %s to repository: %w",
			event.Source(),
			event.Type(),
			err)
	}
	txn.Commit()
	return nil
}

// TODO add context
// TODO use memdb watch
// TODO use source to get the correct memdb
func (r *Repository) Stream(source string) (<-chan mario.CloudEvent, error) {
	ch := make(chan mario.CloudEvent)
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			<-ticker.C
			err := r.getAndSendNonProcessedEvents(ch)
			if err != nil {
				// TODO do something with err
				close(ch)
			}
		}
	}()
	return ch, nil
}

func (r *Repository) getAndSendNonProcessedEvents(ch chan mario.CloudEvent) error {
	txn := r.db.Txn(false)
	resultIter, err := txn.Get("events", "status", StorableEventStatusNotProcessed)
	defer txn.Abort()
	if err != nil {
		return fmt.Errorf("failed getting not processed events: %w", err)
	}
	for obj := resultIter.Next(); obj != nil; obj = resultIter.Next() {
		storableCloudEvent := obj.(StorableCloudEvent)
		cloudEvent, _ := r.cloudEventBuilder.
			Id(storableCloudEvent.ID).
			Source(storableCloudEvent.Source).
			EventType(storableCloudEvent.Type).
			CorrelationID(storableCloudEvent.CorrelationID).
			SpecVersion("").
			Time(storableCloudEvent.Time).
			Data(storableCloudEvent.Data).
			Build()
		ch <- cloudEvent
	}
	return nil
}
