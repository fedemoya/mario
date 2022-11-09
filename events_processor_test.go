package mario

import (
	"testing"
)

func TestEventProcessor(t *testing.T) {

	//wg := &sync.WaitGroup{}
	//wg.Add(2)
	//
	//handlerMock := &HandlerMock[EventMockVisitor]{}
	//handlerMock.On("Start").Return(nil)
	//
	//eventsCh := make(chan AcknowledgeableEvent[EventMockVisitor])
	//handlerMock.On("Subscribe").Return(eventsCh)
	//
	//errorsCh := make(chan error)
	//handlerMock.On("ErrorsCh").Return(errorsCh)
	//
	//visitorMock := &EventMockVisitorMock{}
	//
	//eventMock := &EventMock{}
	//eventMock.On("Accept", visitorMock).Return(nil).Run(
	//	func(args mock.Arguments) {
	//		wg.Done()
	//	},
	//)
	//
	//acknowledgementHandlerMock := &AcknowledgementHandlerMock[EventMockVisitor]{}
	//acknowledgementHandlerMock.On("OnSuccess", eventMock).Return(nil)
	//
	//var processingErr error
	//
	//eventsProcessor := NewProcessor[EventMockVisitor](
	//	handlerMock,
	//	visitorMock,
	//	func(err error) {
	//		processingErr = err
	//		wg.Done()
	//	},
	//	acknowledgementHandlerMock,
	//)
	//
	//eventsProcessor.Start()
	//
	//time.Sleep(1 * time.Millisecond)
	//
	//eventsCh <- eventMock
	//errorsCh <- fmt.Errorf("boom")
	//
	//wg.Wait()
	//
	//eventsProcessor.Stop()
	//
	//eventMock.AssertExpectations(t)
	//acknowledgementHandlerMock.AssertExpectations(t)
	//handlerMock.AssertExpectations(t)
	//
	//assert.Error(t, processingErr)
}
