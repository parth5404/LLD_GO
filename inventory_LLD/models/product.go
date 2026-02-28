package models

type ProductCategory string

const (
	Electronics ProductCategory = "ELECTRONICS"
	Grocery     ProductCategory = "GROCERY"
)

type Product struct {
	sku       string
	name      string
	price     float32
	brand     string
	quantity  int
	threshold int
	category  ProductCategory
}

func NewProduct(sku string, name string, price float32, brand string,
	quantity int, threshold int, category ProductCategory) *Product {
	return &Product{
		sku, name, price,
		brand, quantity, threshold, category,
	}
}
func (p *Product) GetSKU() string {
	return p.sku
}

func (p *Product) GetCategory() ProductCategory {
	return p.category
}

func (p *Product) GetPrice() float32 {
	return p.price
}

func (p *Product) GetQuantity() int {
	return p.quantity
}

func (p *Product) Increment(quantity int) {
	p.quantity += quantity
}

func (p *Product) Decrement(quantity int) {
	p.quantity -= quantity
}
