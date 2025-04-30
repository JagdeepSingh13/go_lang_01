package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// stripePatmentGw := stripe{}
	// razorpayPaymentGw.pay(amount)
	// stripePatmentGw.pay(amount)

	p.gateway.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("Making Payment (Razorpay)...", amount)
}

type stripe struct {
}

func (s stripe) pay(amount float32) {
	fmt.Println("Making Payment (Stripe)...", amount)
}

func main() {
	// stripePatmentGw := stripe{}

	// myPay := payment{gateway: stripePatmentGw}
	// myPay.makePayment(100)

	stripe := stripe{}
	myPay := payment{
		gateway: stripe,
	}
	myPay.makePayment(100)
}
