package main

type CartReceipt interface {
	getFinalAmount() float32
}
type BaseCart struct {
	arr []float32
}

func NewBaseCart(arr []float32) *BaseCart {
	return &BaseCart{
		arr: arr,
	}
}

func (b *BaseCart) getFinalAmount() float32 {
	var sum float32 = 0
	for _, v := range b.arr {
		sum += v
	}
	return sum
}

type EmployeeDiscountDecorator struct {
	Decorator CartReceipt
}

func NewEmployeeDiscountDecorator(decorator CartReceipt) *EmployeeDiscountDecorator {
	return &EmployeeDiscountDecorator{
		Decorator: decorator,
	}
}

func (e *EmployeeDiscountDecorator) getFinalAmount() float32 {
	return e.Decorator.getFinalAmount() * 0.80
}

type FestiveCouponDecorator struct {
	Decorator CartReceipt
}

func NewFestiveCouponDecorator(decorator CartReceipt) *FestiveCouponDecorator {
	return &FestiveCouponDecorator{
		Decorator: decorator,
	}
}

func (f *FestiveCouponDecorator) getFinalAmount() float32 {
	return f.Decorator.getFinalAmount() - 100
}

type TaxDecorator struct {
	Decorator CartReceipt
}

func NewTaxDecorator(decorator CartReceipt) *TaxDecorator {
	return &TaxDecorator{
		Decorator: decorator,
	}
}

func (t *TaxDecorator) getFinalAmount() float32 {
	return t.Decorator.getFinalAmount() * 1.05
}

type SurgeDeliveryDecorator struct {
	Decorator CartReceipt
}

func NewSurgeDeliveryDecorator(decorator CartReceipt) *SurgeDeliveryDecorator {
	return &SurgeDeliveryDecorator{
		Decorator: decorator,
	}
}

func (s *SurgeDeliveryDecorator) getFinalAmount() float32 {
	return s.Decorator.getFinalAmount() + 50
}

// func main() {
// 	arr := []float32{220.00, 780.00}
// 	var cart CartReceipt = NewBaseCart(arr)
// 	cart = NewFestiveCouponDecorator(cart)
// 	cart = NewTaxDecorator(cart)
// 	cart = NewSurgeDeliveryDecorator(cart)
// 	fmt.Println(cart.getFinalAmount())
// }
