package models

type GroceryProduct struct {
	Product
}

func NewGroceryProduct(sku string, name string, price float32, brand string,
	quantity int, threshold int) *GroceryProduct {
	p := Product{
		sku:       sku,
		name:      name,
		price:     price,
		brand:     brand,
		quantity:  quantity,
		threshold: threshold,
		category:  ProductCategory(Grocery),
	}
	return &GroceryProduct{p}
}
