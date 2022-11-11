package events

import (
	"fmt"
	"mario/examples/gateway/domain/paymentapi"
)

type DispatchingVisitor struct {
	paymentapiClient paymentapi.Client
}

func NewDispatchingVisitor(paymentapiClient paymentapi.Client) *DispatchingVisitor {
	return &DispatchingVisitor{paymentapiClient: paymentapiClient}
}

func (d *DispatchingVisitor) VisitDinopayPaymentCreated(created DinopayPaymentCreated) error {
	return d.updateWithdrawal(created.PaymentapiWithdrawalId, created.DinopayStatus)
}

func (d *DispatchingVisitor) VisitDinopayPaymentUpdated(updated DinopayPaymentUpdated) error {
	return d.updateWithdrawal(updated.PaymentapiWithdrawalId, updated.DinopayStatus)
}

func (d *DispatchingVisitor) updateWithdrawal(withdrawalId string, status string) error {
	err := d.paymentapiClient.UpdateWithdrawal(paymentapi.UpdateWithdrawalRequest{
		WithdrawalId: withdrawalId,
		Status:       status,
	})
	if err != nil {
		return fmt.Errorf("failed updating withdrwal on paymentapi: %w", err)
	}
	return nil
}
