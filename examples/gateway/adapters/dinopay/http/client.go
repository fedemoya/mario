package http

import (
	"github.com/google/uuid"
	"mario/examples/gateway/domain/dinopay"
	"time"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (Client) CreatePayment(req dinopay.CreatePaymentRequest) (dinopay.CreatePaymentResponse, error) {
	// TODO implement
	return dinopay.CreatePaymentResponse{
		PaymentId: uuid.New().String(),
		Status:    dinopay.PaymentStatusSubmitted,
		Time:      time.Now().Unix(),
	}, nil
}
