package events

import (
	"mario"
	"mario/examples/gateway/domain"
)

type WithdrawalCreated struct {
	domain.BaseEvent

	Id                 string
	Amount             float64
	SourceAccount      string
	DestinationAccount string
}

var _ mario.Event[Visitor] = WithdrawalCreated{}

func (w WithdrawalCreated) Accept(visitor Visitor) error {
	return visitor.VisitWithdrawalCreated(w)
}
