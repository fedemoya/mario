package models

const (
	QuotePayoutTypeBankTransfer string = "BANK_TRANSFER"
	QuotePayoutTypeBalance             = "BALANCE"
	QuotePayoutTypeSwift               = "SWIFT"
	QuotePayoutTypeSwiftOur            = "SWIFT_OUR"
	QuotePayoutTypeInterac             = "INTERAC"
)

type CreateQuoteRequestBody struct {
	SourceCurrency *string
	TargetCurrency *string
	SourceAmount   *float64
	TargetAmount   *float64
	PayOut         *string
	PreferredPayIn *string
}
