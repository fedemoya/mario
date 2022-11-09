package paymentapi

type UpdateWithdrawalRequest struct {
	WithdrawalId string
	Status       string
}

type Client interface {
	UpdateWithdrawal(UpdateWithdrawalRequest) error
}
