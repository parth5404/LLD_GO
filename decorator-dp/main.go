package main

import "fmt"

func main() {
	cfitemfactory := NewCoffeeFactory()
	notifier := NewPublisher()
	notifier.AddObserver(
		NewCustomer("Alice"),
		NewKitchenDisplay("K1"),
		NewInventorySystem(),
	)
	paymentfactory := NewPaymentFactory(notifier)
	
	finalitem, err := cfitemfactory.GetBaseCoffee("LATTE")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	finalitem, err = cfitemfactory.AddExtra("MILK", finalitem)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	paymentStrategy, err := paymentfactory.GetPaymentInterface("UPI")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	paymentStrategy.pay(finalitem)
}
