package models

import "github.com/google/uuid"

type TransferDetailsBuilder struct {
	transferDetails *TransferDetails
}

func NewTransferDetailsBuilder() *TransferDetailsBuilder {
	transferDetails := &TransferDetails{}
	b := &TransferDetailsBuilder{transferDetails: transferDetails}
	return b
}

func (b *TransferDetailsBuilder) Reference(reference string) *TransferDetailsBuilder {
	b.transferDetails.Reference = reference
	return b
}

func (b *TransferDetailsBuilder) TransferPurpose(transferPurpose string) *TransferDetailsBuilder {
	b.transferDetails.TransferPurpose = transferPurpose
	return b
}

func (b *TransferDetailsBuilder) SourceOfFunds(sourceOfFunds string) *TransferDetailsBuilder {
	b.transferDetails.SourceOfFunds = sourceOfFunds
	return b
}

func (b *TransferDetailsBuilder) Build() TransferDetails {
	return *b.transferDetails
}

type CreateTransferRequestBodyBuilder struct {
	createTransferRequestBody *CreateTransferRequestBody
}

func NewCreateTransferRequestBodyBuilder() *CreateTransferRequestBodyBuilder {
	createTransferRequestBody := &CreateTransferRequestBody{}
	b := &CreateTransferRequestBodyBuilder{createTransferRequestBody: createTransferRequestBody}
	return b
}

func (b *CreateTransferRequestBodyBuilder) TargetAccount(targetAccount int64) *CreateTransferRequestBodyBuilder {
	b.createTransferRequestBody.TargetAccount = &targetAccount
	return b
}

func (b *CreateTransferRequestBodyBuilder) QuoteUuid(quoteUuid uuid.UUID) *CreateTransferRequestBodyBuilder {
	b.createTransferRequestBody.QuoteUuid = &quoteUuid
	return b
}

func (b *CreateTransferRequestBodyBuilder) CustomerTransactionId(customerTransactionId uuid.UUID) *CreateTransferRequestBodyBuilder {
	b.createTransferRequestBody.CustomerTransactionId = &customerTransactionId
	return b
}

func (b *CreateTransferRequestBodyBuilder) Details(details TransferDetails) *CreateTransferRequestBodyBuilder {
	b.createTransferRequestBody.Details = &details
	return b
}

func (b *CreateTransferRequestBodyBuilder) Build() CreateTransferRequestBody {
	return *b.createTransferRequestBody
}
