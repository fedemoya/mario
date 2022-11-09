package events

type Visitor interface {
	VisitPaymentStatusUpdated(PaymentStatusUpdated) error
}
