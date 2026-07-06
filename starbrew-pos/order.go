package main

import (
	"fmt"
	"sync"
	"time"
)

type OrderStatus int

const (
	Pending OrderStatus = iota
	Preparing
	Ready
	Completed
	Cancelled
)

func (s OrderStatus) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Preparing:
		return "Preparing"
	case Ready:
		return "Ready"
	case Completed:
		return "Completed"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}

type OrderService struct {
	publisher      *Publisher
	paymentFactory *PaymentFactory
}

func NewOrderService(pub *Publisher, pf *PaymentFactory) *OrderService {
	return &OrderService{
		publisher:      pub,
		paymentFactory: pf,
	}
}

type Order struct {
	OrderStatus   OrderStatus
	PaymentStatus PaymentStatus
	CreatedAt     time.Time
	CancelUntil   time.Time
	Product       CoffeeItem
	mu            sync.RWMutex
}

func (os *OrderService) CreateOrder(product CoffeeItem) *Order {
	now := time.Now()
	order := &Order{
		OrderStatus:   Pending,
		CreatedAt:     now,
		CancelUntil:   now.Add(5 * time.Minute),
		Product:       product,
		PaymentStatus: PendingP,
	}
	fmt.Println("Creating the Order", product.CoffeeItem())
	os.publisher.NotifyAll(OrderReceivedEvent{Order: order})
	return order
}

func (os *OrderService) PayOrder(payType string, order *Order) (*Order, error) {
	order.mu.Lock()
	if order.PaymentStatus == PaidP {
		order.mu.Unlock()
		return nil, fmt.Errorf("Order is already paid!")
	}
	if order.PaymentStatus == ProcessingPaymentP {
		order.mu.Unlock()
		return nil, fmt.Errorf("Payment is already in progress, please wait!")
	}
	if order.OrderStatus == Cancelled {
		order.mu.Unlock()
		return nil, fmt.Errorf("Cannot pay for cancelled order!")
	}
	order.PaymentStatus = ProcessingPaymentP
	order.mu.Unlock()

	start, _ := os.paymentFactory.GetPaymentInterface(payType)
	start.pay(order)

	order.mu.Lock()
	if payType != "COD" {
		order.PaymentStatus = PaidP
	}
	order.mu.Unlock()

	os.publisher.NotifyAll(PaymentProcessedEvent{Order: order})

	go os.processOrder(order)

	return order, nil
}

func (os *OrderService) CancelOrder(o *Order) error {
	o.mu.Lock()
	defer o.mu.Unlock()
	
	if time.Now().After(o.CancelUntil) {
		return fmt.Errorf("Deadline over to cancel order")
	}
	if o.OrderStatus == Completed {
		return fmt.Errorf("Order is already completed")
	}
	o.OrderStatus = Cancelled
	return nil
}

func (os *OrderService) processOrder(o *Order) {
	os.publisher.NotifyAll(OrderProcessingEvent{Order: o})

	time.Sleep(5 * time.Second) // Coffee ban rahi hai...

	o.mu.Lock()
	if o.OrderStatus == Cancelled {
		o.mu.Unlock()
		return
	}
	o.OrderStatus = Completed
	o.mu.Unlock()

	os.publisher.NotifyAll(OrderCompletedEvent{Order: o})
}
