package models

type ElectronicsProduct struct {
	Product
}

func NewElectronicsProduct(sku string, name string, price float32, brand string,
	quantity int, threshold int) *ElectronicsProduct {
	p := Product{
		sku:       sku,
		name:      name,
		price:     price,
		brand:     brand,
		quantity:  quantity,
		threshold: threshold,
		category:  ProductCategory(Electronics),
	}
	return &ElectronicsProduct{p}
}
