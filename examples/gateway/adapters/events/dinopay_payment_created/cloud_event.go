package dinopay_payment_created

type cloudEvent struct {
	id             string
	source         string
	cloudEventType string
	time           int64
	correlationID  string
}

func (c cloudEvent) ID() string {
	return c.id
}

func (c cloudEvent) Source() string {
	return c.source
}

func (c cloudEvent) Type() string {
	return c.cloudEventType
}

func (c cloudEvent) Time() int64 {
	return c.time
}

func (c cloudEvent) CorrelationID() string {
	return c.correlationID
}
