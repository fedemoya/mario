package dinopay_payment_created

import (
	"encoding/json"
	"mario/examples/gateway/domain/events"
)

type dinopayEventCreatedWithTags struct {
	PaymentapiWithdrawalId string `json:"paymentapi_withdrawal_id"`
	DinopayId              string `json:"dinopay_id"`
	DinopayStatus          string `json:"dinopay_status"`
	DinopayTime            int64  `json:"dinopay_time"`
}

type jsonMarshaler struct {
	event events.DinopayPaymentCreated
}

func (jm jsonMarshaler) MarshalJSON() ([]byte, error) {
	dinopayEventCreatedWithTags := dinopayEventCreatedWithTags{
		PaymentapiWithdrawalId: jm.event.PaymentapiWithdrawalId,
		DinopayId:              jm.event.DinopayId,
		DinopayStatus:          jm.event.DinopayStatus,
		DinopayTime:            jm.event.DinopayTime,
	}
	return json.Marshal(dinopayEventCreatedWithTags)
}
