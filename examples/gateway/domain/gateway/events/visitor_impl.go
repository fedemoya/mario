package events

import (
	"fmt"
	"github.com/rs/zerolog"
	"mario/examples/gateway/domain/paymentapi"
	"os"
)

type VisitorImpl struct {
	paymentapiClient paymentapi.Client
	logger           zerolog.Logger
}

func NewVisitorImpl(paymentapiClient paymentapi.Client) *VisitorImpl {
	return &VisitorImpl{
		paymentapiClient: paymentapiClient,
		logger:           zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

func (d *VisitorImpl) VisitDinopayPaymentCreated(created DinopayPaymentCreated) error {
	d.logger.Info().
		Str("paymentapiWithdrawalId", created.PaymentapiWithdrawalId).
		Msgf("visiting %s event", created.Type())
	return d.updateWithdrawal(created.PaymentapiWithdrawalId, created.DinopayStatus)
}

func (d *VisitorImpl) VisitDinopayPaymentUpdated(updated DinopayPaymentUpdated) error {
	d.logger.Info().
		Str("paymentapiWithdrawalId", updated.PaymentapiWithdrawalId).
		Msgf("visiting %s event", updated.Type())
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
