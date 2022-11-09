package events

type WithdrawalCreated struct {
	Id                 string  `json:"id"`
	Amount             float64 `json:"amount"`
	SourceAccount      string  `json:"source_account"`
	DestinationAccount string  `json:"destination_account"`
}
