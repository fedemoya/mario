package events

import (
	"fmt"
	"mario/examples/gateway/domain/paymentapi"
)

type VisitorImpl struct {
	paymentapiClient paymentapi.Client
}

func NewVisitorImpl(paymentapiClient paymentapi.Client) *VisitorImpl {
	return &VisitorImpl{paymentapiClient: paymentapiClient}
}

func (d *VisitorImpl) VisitDinopayPaymentCreated(created DinopayPaymentCreated) error {
	return d.updateWithdrawal(created.PaymentapiWithdrawalId, created.DinopayStatus)
}

func (d *VisitorImpl) VisitDinopayPaymentUpdated(updated DinopayPaymentUpdated) error {
	return d.updateWithdrawal(updated.PaymentapiWithdrawalId, updated.DinopayStatus)
}

func (d *VisitorImpl) updateWithdrawal(withdrawalId string, status string) error {
	err := d.paymentapiClient.UpdateWithdrawal(paymentapi.UpdateWithdrawalRequest{
		WithdrawalId: withdrawalId,
		Status:       status,
	})
	if err != nil {
		return fmt.Errorf("failed updating withdrwal on paymentapi: %w", err)
	}
	return nil
}
