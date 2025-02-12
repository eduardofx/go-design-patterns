package main

import (
	"errors"
	"fmt"
)

type Payment interface {
	Pay(amount float64) string
}

type CreditCard struct{}
type Paypal struct{}

func (c CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid $ %.2f useing Credit Card", amount)
}

func (p Paypal) Pay(amount float64) string {
	return fmt.Sprintf("Paid $ %.2f using Paypal", amount)
}

var PaymentFactory = map[string]func() Payment{
	"credit_card": func() Payment { return CreditCard{} },
	"paypal":      func() Payment { return Paypal{} },
}

func NewPayment(method string) (Payment, error) {
	if factory, exists := PaymentFactory[method]; exists {
		return factory(), nil
	}
	return nil, errors.New("invalid payment method")
}

func main() {
	payment, err := NewPayment("credit_card")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(payment.Pay(100.00))
}
