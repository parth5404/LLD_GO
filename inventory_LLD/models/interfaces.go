package models

type InventoryItem interface {
	GetSKU() string
	GetCategory() ProductCategory
	GetPrice() float32
	GetQuantity() int
	Increment(quantity int)
	Decrement(quantity int)
}

type ReplenishmentStrategy interface {
	Replenish(product InventoryItem)
}

type Observer interface {
	Update()
}

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}
