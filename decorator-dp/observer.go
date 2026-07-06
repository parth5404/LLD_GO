package main

import (
	"fmt"
)

type Event interface {
	EventName() string
}

type OrderReceivedEvent struct {
	*Order
}

func (e OrderReceivedEvent) EventName() string {
	return "Order Received"
}

type PaymentProcessedEvent struct {
	*Order
}

func (e PaymentProcessedEvent) EventName() string {
	return "Payment Processed"
}

type OrderProcessingEvent struct {
	*Order
	ItemName string
}

func (e OrderProcessingEvent) EventName() string {
	return "Order is Processing " + e.ItemName
}

type OrderCompletedEvent struct {
	*Order
}

func (e OrderCompletedEvent) EventName() string {
	return "Order is Made"
}

type Observer interface {
	Update(e Event)
}

type Customer struct {
	Name string
}

func NewCustomer(name string) *Customer {
	return &Customer{Name: name}
}

func (c *Customer) Update(e Event) {
	switch event := e.(type) {
	case OrderReceivedEvent:
		event.mu.RLock() // Safe Read Lock
		fmt.Printf("[Customer %s] Notification: %s | Order Status: %s\n", c.Name, event.EventName(), event.OrderStatus)
		event.mu.RUnlock()
	case PaymentProcessedEvent:
		event.mu.RLock()
		fmt.Printf("[Customer %s] Notification: %s | Payment Status: %s\n", c.Name, event.EventName(), event.PaymentStatus)
		event.mu.RUnlock()
	case OrderProcessingEvent:
		event.mu.RLock()
		fmt.Printf("[Customer %s] Notification: %s | Order Status: %s\n", c.Name, event.EventName(), event.OrderStatus)
		event.mu.RUnlock()
	case OrderCompletedEvent:
		event.mu.RLock()
		fmt.Printf("[Customer %s] Notification: %s | Order Status: %s\n", c.Name, event.EventName(), event.OrderStatus)
		event.mu.RUnlock()
	}
}

type KitchenDisplay struct {
	ID string
}

func NewKitchenDisplay(id string) *KitchenDisplay {
	return &KitchenDisplay{ID: id}
}

func (k *KitchenDisplay) Update(e Event) {
	switch e.(type) {
	case OrderReceivedEvent, OrderProcessingEvent:
		// Since we handle two events here, we typecast them back to access the embedded *Order
		var order *Order
		var eventName string

		if ev, ok := e.(OrderReceivedEvent); ok {
			order = ev.Order
			eventName = ev.EventName()
		} else if ev, ok := e.(OrderProcessingEvent); ok {
			order = ev.Order
			eventName = ev.EventName()
		}

		order.mu.RLock()
		fmt.Printf("[Kitchen Display %s] Notification: %s | Status: %s\n", k.ID, eventName, order.OrderStatus)
		order.mu.RUnlock()
	}
}

type InventorySystem struct {
}

func NewInventorySystem() *InventorySystem {
	return &InventorySystem{}
}

func (i *InventorySystem) Update(e Event) {
	switch event := e.(type) {
	case OrderCompletedEvent:
		event.mu.RLock()
		fmt.Printf("[Inventory System] Notification: %s (Deducting items) | Final Status: %s\n", event.EventName(), event.OrderStatus)
		event.mu.RUnlock()
	}
}

type Publisher struct {
	observers []Observer
}

func NewPublisher() *Publisher {
	return &Publisher{
		observers: make([]Observer, 0),
	}
}

func (p *Publisher) AddObserver(o ...Observer) {
	p.observers = append(p.observers, o...)
}

func (p *Publisher) NotifyAll(e Event) {
	for _, observer := range p.observers {
		observer.Update(e)
	}
}
