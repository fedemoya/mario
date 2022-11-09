package events

type Visitor interface {
	VisitWithdrawalCreated(withdrawCreated WithdrawalCreated) error
}
