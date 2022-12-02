package events

type Visitor interface {
	VisitDinopayPaymentCreated(DinopayPaymentCreated) error
	VisitDinopayPaymentUpdated(DinopayPaymentUpdated) error
}
