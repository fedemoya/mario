package paymentapi

type EventVisitor interface {
	VisitWithdrawalCreated(wCreated WithdrawalCreated) error
}
