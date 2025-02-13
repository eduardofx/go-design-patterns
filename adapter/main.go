package main

import "fmt"

type PaymentProcessor interface {
	Pay(amount float64)
}

type LegacyPaymentGateway struct{}

type PaymentAdapter struct {
	legacy *LegacyPaymentGateway
}

func (l *LegacyPaymentGateway) ProcessorPayment(value float64) {
	fmt.Printf("Processing payment of $%.2f through Legacy Payment Gateway\n", value)
}

func (p *PaymentAdapter) Pay(amount float64) {
	p.legacy.ProcessorPayment(amount)
}

func main() {
	legacyGateway := &LegacyPaymentGateway{}
	adapter := &PaymentAdapter{legacyGateway}

	adapter.Pay(100.20)
}
