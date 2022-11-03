package mario

type AcknowledgementHandler interface {
	OnSuccess(acknowledger Acknowledger) error
	OnError(acknowledger Acknowledger, err error) error
}
