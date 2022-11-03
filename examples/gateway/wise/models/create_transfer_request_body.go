package models

import (
	"github.com/google/uuid"
)

type TransferDetails struct {
	Reference       string
	TransferPurpose string
	SourceOfFunds   string
}

type CreateTransferRequestBody struct {
	TargetAccount         *int64
	QuoteUuid             *uuid.UUID
	CustomerTransactionId *uuid.UUID
	Details               *TransferDetails
}
