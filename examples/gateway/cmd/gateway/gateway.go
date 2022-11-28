package main

func main() {

	//gatewayDomainEventsVisitor := gatewayDomainEvents.NewVisitorImpl(paymentapiHttp.NewClient())
	//
	//paymentapiEventsVisitor := paymentapiDomainEvents.NewVisitorImpl(
	//	dinopayHttp.NewClient(),
	//	gatewayDomainEventsVisitor,
	//)
	//
	//paymentApiEventsSource := amqp.NewEventsSource()
	//paymentApiEventsFactory := events.NewFactory()
	//
	//paymentApiEventsProcessor := mario.NewProcessor[paymentapiDomainEvents.Visitor](
	//	paymentApiEventsSource,
	//	paymentApiEventsFactory,
	//	paymentapiEventsVisitor,
	//	func(err error) {
	//		fmt.Printf("paymentapi events processor error: %s\n", err.Error())
	//	},
	//)
	//
	//paymentApiEventsProcessor.Start()
	//
	//fmt.Println("Gateway started")
	//
	//time.Sleep(10 * time.Minute)
}
