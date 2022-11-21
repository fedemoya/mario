package dinopay_payment_created

import (
	"encoding/json"
	gatewayDomainEvents "mario/examples/gateway/domain/events"
)

type dinopayEventCreatedWithTags struct {
	PaymentapiWithdrawalId string `json:"paymentapi_withdrawal_id"`
	DinopayId              string `json:"dinopay_id"`
	DinopayStatus          string `json:"dinopay_status"`
	DinopayTime            int64  `json:"dinopay_time"`
}

func marshalJSON(event gatewayDomainEvents.DinopayPaymentCreated) ([]byte, error) {
	dinopayEventCreatedWithTags := dinopayEventCreatedWithTags{
		PaymentapiWithdrawalId: event.PaymentapiWithdrawalId,
		DinopayId:              event.DinopayId,
		DinopayStatus:          event.DinopayStatus,
		DinopayTime:            event.DinopayTime,
	}
	return json.Marshal(dinopayEventCreatedWithTags)
}
