package http

import (
	"github.com/rs/zerolog"
	"mario/examples/gateway/domain/paymentapi"
	"os"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) UpdateWithdrawal(req paymentapi.UpdateWithdrawalRequest) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Info().Msgf("sending update withdrawal req to payment api")
	return nil
}
