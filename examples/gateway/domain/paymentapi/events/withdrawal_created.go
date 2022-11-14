package events

import (
	"mario"
)

type WithdrawalCreated struct {
	mario.BaseEvent

	Id                 string
	Amount             float64
	SourceAccount      string
	DestinationAccount string
}

var _ mario.Event[Visitor] = WithdrawalCreated{}

func (w WithdrawalCreated) Accept(visitor Visitor) error {
	return visitor.VisitWithdrawalCreated(w)
}
