package main

import "fmt"

type PaymentStatus int

const (
	PendingP PaymentStatus = iota
	ProcessingPaymentP
	PaidP
	CancelledP
)

type PaymentFactory struct {
}

func NewPaymentFactory() *PaymentFactory {
	return &PaymentFactory{}
}

type PaymentStrategy interface {
	pay(order *Order) int
}

type UPI struct {
}

func (u *UPI) pay(order *Order) int {
	price := order.Product.price()
	return price
}

type Wallet struct {
}

func (w *Wallet) pay(order *Order) int {
	price := order.Product.price()
	return price
}

type CreditCard struct {
}

func (c *CreditCard) pay(order *Order) int {
	price := order.Product.price()
	return price
}

func (pf *PaymentFactory) GetPaymentInterface(types string) (PaymentStrategy, error) {
	switch types {
	case "UPI":
		return &UPI{}, nil
	case "WALLET":
		return &Wallet{}, nil
	case "CREDIT_CARD":
		return &CreditCard{}, nil
	default:
		return nil, fmt.Errorf("unknown payment method: %s", types)
	}
}
