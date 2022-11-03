package withdrawal

import (
	"context"
	papimodels "mario/examples/gateway/paymentapi/models"
	"mario/examples/gateway/wise"
	"mario/examples/gateway/wise/models"
)

type Service struct {
	wiseClient wise.IHTTPClient
}

func (s *Service) CreateWiseQuote(profileID int64, withdrawal papimodels.Withdrawal) {
	createQuoteReq := models.NewCreateQuoteRequestBodyBuilder().
		PayOut(models.QuotePayoutTypeBalance).
		SourceAmount(withdrawal.Amount).
		SourceCurrency(withdrawal.Currency).
		TargetCurrency(withdrawal.Beneficiary.Currency).
		Build()

	res, err := s.wiseClient.CreateQuote(context.TODO(), profileID, createQuoteReq)
}
