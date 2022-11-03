package mario

import (
	"fmt"
	"sync"
)

type EventBus interface {
	Subscribe(topic string) <-chan any
	Unsubscribe(topic string, ch <-chan any) error
	Publish(topic string, event any) error
}

var _ EventBus = (*EventBusImpl)(nil)

type EventBusImpl struct {
	subscribers       map[string][]chan any
	channelBufferSize int
	sync.Mutex
}

func NewPublisher(bufferSize int) *EventBusImpl {
	return &EventBusImpl{
		channelBufferSize: bufferSize,
		subscribers:       make(map[string][]chan any, 0),
	}
}

func (s *EventBusImpl) Subscribe(topic string) <-chan any {

	ch := make(chan any, s.channelBufferSize)

	s.Lock()
	defer s.Unlock()

	topicSubscribers, ok := s.subscribers[topic]
	if !ok {
		topicSubscribers = make([]chan any, 0)
		s.subscribers[topic] = topicSubscribers
	}
	topicSubscribers = append(topicSubscribers, ch)

	return ch
}

func (s *EventBusImpl) Unsubscribe(topic string, c <-chan any) error {

	s.Lock()
	defer s.Unlock()

	topicSubscribers, ok := s.subscribers[topic]
	if !ok {
		return fmt.Errorf("topic %s doesn't exist", topic)
	}

	length := len(topicSubscribers)
	for i := 0; i < length; i++ {
		if topicSubscribers[i] == c {
			close(topicSubscribers[i])
			topicSubscribers[i] = topicSubscribers[length-1]
			topicSubscribers = topicSubscribers[:length-1]
			return nil
		}
	}

	return nil
}

func (s *EventBusImpl) Publish(topic string, event any) error {

	topicSubscribers, ok := s.subscribers[topic]
	if !ok {
		return fmt.Errorf("topic %s doesn't exist", topic)
	}

	for _, ch := range topicSubscribers {
		ch <- event
	}

	return nil
}

func (s *EventBusImpl) UnsubscribeAll() {

	s.Lock()
	defer s.Unlock()

	for _, topicSubscribers := range s.subscribers {
		for _, ch := range topicSubscribers {
			close(ch)
		}
	}
}
