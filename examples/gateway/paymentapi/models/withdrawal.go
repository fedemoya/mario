package models

type Beneficiary struct {
	AccountId string
	Currency  string
}

type Withdrawal struct {
	WithdrawalId string
	Currency     string
	Amount       float64
	Beneficiary  Beneficiary
}
