package main

import "fmt"

type Observer interface {
	Update(message string)
}

type Customer struct {
	Name string
}

func NewCustomer(name string) *Customer {
	return &Customer{Name: name}
}

func (c *Customer) Update(message string) {
	fmt.Printf("[Customer %s] Notification: %s\n", c.Name, message)
}

type KitchenDisplay struct {
	ID string
}

func NewKitchenDisplay(id string) *KitchenDisplay {
	return &KitchenDisplay{ID: id}
}

func (k *KitchenDisplay) Update(message string) {
	fmt.Printf("[Kitchen Display %s] Notification: %s\n", k.ID, message)
}

type InventorySystem struct {
}

func NewInventorySystem() *InventorySystem {
	return &InventorySystem{}
}

func (i *InventorySystem) Update(message string) {
	fmt.Printf("[Inventory System] Notification: %s\n", message)
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

func (p *Publisher) NotifyAll(message string) {
	for _, observer := range p.observers {
		observer.Update(message)
	}
}
