package models

type CreateQuoteRequestBodyBuilder struct {
	createQuoteRequestBody *CreateQuoteRequestBody
}

func NewCreateQuoteRequestBodyBuilder() *CreateQuoteRequestBodyBuilder {
	createQuoteRequestBody := &CreateQuoteRequestBody{}
	b := &CreateQuoteRequestBodyBuilder{createQuoteRequestBody: createQuoteRequestBody}
	return b
}

func (b *CreateQuoteRequestBodyBuilder) SourceCurrency(sourceCurrency string) *CreateQuoteRequestBodyBuilder {
	b.createQuoteRequestBody.SourceCurrency = &sourceCurrency
	return b
}

func (b *CreateQuoteRequestBodyBuilder) TargetCurrency(targetCurrency string) *CreateQuoteRequestBodyBuilder {
	b.createQuoteRequestBody.TargetCurrency = &targetCurrency
	return b
}

func (b *CreateQuoteRequestBodyBuilder) SourceAmount(sourceAmount float64) *CreateQuoteRequestBodyBuilder {
	b.createQuoteRequestBody.SourceAmount = &sourceAmount
	return b
}

func (b *CreateQuoteRequestBodyBuilder) TargetAmount(targetAmount float64) *CreateQuoteRequestBodyBuilder {
	b.createQuoteRequestBody.TargetAmount = &targetAmount
	return b
}

func (b *CreateQuoteRequestBodyBuilder) PayOut(payOut string) *CreateQuoteRequestBodyBuilder {
	b.createQuoteRequestBody.PayOut = &payOut
	return b
}

func (b *CreateQuoteRequestBodyBuilder) PreferredPayIn(preferredPayIn *string) *CreateQuoteRequestBodyBuilder {
	b.createQuoteRequestBody.PreferredPayIn = preferredPayIn
	return b
}

func (b *CreateQuoteRequestBodyBuilder) Build() CreateQuoteRequestBody {
	return *b.createQuoteRequestBody
}
