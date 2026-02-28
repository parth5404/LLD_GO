package factory

import "inventory_LLD/models"

type ProductFactory struct{}

func NewProductFactory() *ProductFactory {
	return &ProductFactory{}
}

func (f *ProductFactory) Create(category models.ProductCategory,
	sku string,
	name string,
	price float32,
	brand string,
	quantity int,
	threshold int) models.InventoryItem {

	switch category {
	case models.Grocery:
		return models.NewGroceryProduct(sku, name, price, brand, quantity, threshold)
	case models.Electronics:
		return models.NewElectronicsProduct(sku, name, price, brand, quantity, threshold)
	default:
		return nil
	}
}
