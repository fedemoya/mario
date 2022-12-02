package events

type BuildersFactory interface {
	CreateDinopayPaymentCreatedBuilder() DinopayPaymentCreatedBuilder
	CreateDinopayPaymentUpdatedBuilder() DinopayPaymentUpdatedBuilder
}
