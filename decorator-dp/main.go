package main

import (
	"fmt"
	"time"
)

func main() {
	cfitemfactory := NewCoffeeFactory()

	publisher := NewPublisher()
	publisher.AddObserver(
		NewCustomer("Alice"),
		NewKitchenDisplay("K1"),
		NewInventorySystem(),
	)

	paymentFactory := NewPaymentFactory()
	orderService := NewOrderService(publisher, paymentFactory)

	finalitem, err := cfitemfactory.GetBaseCoffee("LATTE")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	addons := []string{
		"milk",
		// "chocolate",
		"caramel",
	}

	for _, addon := range addons {
		modifier, err := GetAddon(addon)
		if err != nil {
			panic(err)
		}
		finalitem = modifier(finalitem)
	}

	// paymentStrategy, err := paymentfactory.GetPaymentInterface("UPI")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// paymentStrategy.pay(finalitem)
	order := orderService.CreateOrder(finalitem)
	orderService.PayOrder("UPI", order)
	time.Sleep(7 * time.Second)
}
