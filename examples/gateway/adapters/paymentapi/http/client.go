package http

import (
	"mario/examples/gateway/domain/paymentapi"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) UpdateWithdrawal(_ paymentapi.UpdateWithdrawalRequest) error {
	// TODO implement
	return nil
}
