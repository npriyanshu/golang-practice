package main

import "fmt"

// convention for interfaces is to start the name of the interface with I or end it with er

type paymenter interface {
	// define the method signature that we want to implement
	pay(amount float32)
}

type payment struct {
// we will not give the type of gateway because it can be any type that implements the pay method
// gateway razorpay

// we will use an interface to define the type of gateway
	gateway paymenter
}

func (p payment) pay(amount float32) {
	// fmt.Println("Payment processed")
	p.gateway.pay(amount)
}


type razorpay struct {}

func (r razorpay) pay(amount float32){
	fmt.Println("Payment processed by razorpay $", amount)
}


type stripe struct {}

func (s stripe) pay(amount float32){
	fmt.Println("Payment processed by stripe $", amount)
}

func (s stripe) refund(){
	fmt.Println("Refund processed by stripe")
}

func main() {
razorpaypay := payment{gateway: razorpay{}}
razorpaypay.pay(100.5)

// we cannot use stripepay because it is not of type payment and does not have the pay method
// but if we want to use stripepay we can use an interface
// an interface is a collection of method signatures that a type can implement
    stripeGw := stripe{}
	stripepay := payment{gateway: stripeGw}
	stripepay.pay(200.75)

	// stripepay.gateway.refund() // will not work because gateway is of type paymenter and does not have the refund method

	stripepay.gateway.(stripe).refund() // will work because we are type asserting the gateway to be of type stripe and then calling the refund method
	// type assertion is a way to convert an interface type to a concrete type

	stripeGw.refund() // will work because stripeGw is of type stripe and has the refund method
}	