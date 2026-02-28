package models

type InventoryItem interface {
	GetSKU() string
	GetCategory() ProductCategory
	GetPrice() float32
	GetQuantity() int
	Increment(quantity int)
	Decrement(quantity int)
}
