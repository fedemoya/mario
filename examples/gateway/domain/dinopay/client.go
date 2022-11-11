package dinopay

const (
	PaymentStatusSubmitted = "submitted"
	PaymentStatusFailed    = "failed"
	PaymentStatusSucceed   = "succeed"
)

type CreatePaymentRequest struct {
	SourceAccount      string
	DestinationAccount string
	Amount             float64
	ClientID           string
}

type CreatePaymentResponse struct {
	PaymentId string
	Status    string
	Time      int64
}

type Client interface {
	CreatePayment(req CreatePaymentRequest) (CreatePaymentResponse, error)
}
