package gateway

type DinoPayOutboundPaymentCreated struct {
	DinoPayId    string
	WithdrawalId string
}

func (event DinoPayOutboundPaymentCreated) Accept(visitor EventVisitor) error {
	return visitor.VisitDinoPayOutboundPaymentCreated(event)
}

type DinoPayOutboundPaymentFailed struct {
	WithdrawalId string
	ErrorCode    string
	ErrorMsg     string
}

func (event DinoPayOutboundPaymentFailed) Accept(visitor EventVisitor) error {
	return visitor.VisitDinoPayOutboundPaymentFailed(event)
}

type DinoPayInboundPaymentCreated struct {
	DinoPayId       string
	Currency        string
	Amount          string
	BeneficiaryData string
}

func (event DinoPayInboundPaymentCreated) Accept(visitor EventVisitor) error {
	return visitor.VisitDinoPayInboundPaymentCreated(event)
}
