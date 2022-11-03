package wise

import (
	"context"
	"mario/examples/gateway/wise/models"
)

type IHTTPClient interface {
	CreateQuote(ctx context.Context, profileId int64, createQuote models.CreateQuoteRequestBody) (models.CreateQuoteResponse, error)
}
