package models

import "fmt"

type Warehouse struct {
	id       int
	name     string
	location string
	products map[string]InventoryItem
}

func NewWarehouse(id int, name string, location string) *Warehouse {
	return &Warehouse{
		id:       id,
		name:     name,
		location: location,
		products: make(map[string]InventoryItem),
	}
}

func (w *Warehouse) AddProduct(item InventoryItem) {
	w.products[item.GetSKU()] = item
}

func (w *Warehouse) AddQuantity(sku string, quantity int) {
	item, ok := w.products[sku]
	if !ok {
		fmt.Println("No product with this SKU")
		return
	}
	item.Increment(quantity)
}

func (w *Warehouse) removeProduct(sku string, quantity int) {
	item, ok := w.products[sku]
	if !ok {
		fmt.Println("No product with this SKU")
		return
	}
	item.Decrement(quantity)
	if item.GetQuantity() <= 0 {
		delete(w.products, sku)
	}
}
