package main

import "fmt"

type PaymentFactory struct {
	notifier *Publisher
}

func NewPaymentFactory(notifier *Publisher) *PaymentFactory {
	return &PaymentFactory{
		notifier: notifier,
	}
}

type PaymentStrategy interface {
	pay(final CoffeeItem) int
}

type UPI struct {
	notifier *Publisher
}

func (u *UPI) pay(final CoffeeItem) int {
	price := final.price()
	u.notifier.NotifyAll(fmt.Sprintf("Paid $%d via UPI for [%s]\n", price, final.CoffeeItem()))
	return price
}

type Wallet struct {
	notifier *Publisher
}

func (w *Wallet) pay(final CoffeeItem) int {
	price := final.price()
	w.notifier.NotifyAll(fmt.Sprintf("Paid $%d via Wallet for [%s]\n", price, final.CoffeeItem()))
	return price
}

type CreditCard struct {
	notifier *Publisher
}

func (c *CreditCard) pay(final CoffeeItem) int {
	price := final.price()
	c.notifier.NotifyAll(fmt.Sprintf("Paid $%d via Credit Card for [%s]\n", price, final.CoffeeItem()))
	return price
}

func (pf *PaymentFactory) GetPaymentInterface(types string) (PaymentStrategy, error) {
	switch types {
	case "UPI":
		return &UPI{
			notifier: pf.notifier,
		}, nil
	case "WALLET":
		return &Wallet{
			notifier: pf.notifier,
		}, nil
	case "CREDIT_CARD":
		return &CreditCard{
			notifier: pf.notifier,
		}, nil
	default:
		return nil, fmt.Errorf("unknown payment method: %s", types)
	}
}
