package main

import (
	"inventory_LLD/factory"
	"inventory_LLD/models"
	"inventory_LLD/observer"
	"inventory_LLD/strategy"
)

func main() {
	productfact := factory.NewProductFactory()
	electronicsprod := productfact.Create(models.Electronics, "SKU123", "Laptop",
		32.50, "Dell", 25, 10)
	warehouse1 := models.NewWarehouse(1, "w1", "pune")
	warehouse1.AddProduct(electronicsprod)
	observer1 := observer.NewAdminObserver(0, "sher")
	//observer2 := &observer.EmailObserver{}
	i_manager := models.GetInstance()
	i_manager.AddWarehouse(warehouse1)
	i_manager.RegisterObserver(observer1)
	//i_manager.RemoveObserver(observer2)
	strat := strategy.BulkReplenishment{}
	i_manager.SetReplenishmentStrategy(&strat)
	//i_manager.NotifyObservers()
	i_manager.ReplenishGoods(electronicsprod)

}
