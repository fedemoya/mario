package models

import "time"

type CreateQuoteResponse struct {
	Id                 string
	ExpirationTime     time.Time
	SourceAmount       float64
	Status             string
	Rate               float64
	CreatedTime        time.Time
	RateType           string
	RateExpirationTime time.Time
}
