package main

type Milk struct {
	coffee CoffeeItem
}

func NewMilk(c CoffeeItem) *Milk {
	return &Milk{coffee: c}
}

func (m *Milk) CoffeeItem() string {
	return m.coffee.CoffeeItem() + " + Milk"
}

func (m *Milk) price() int {
	return m.coffee.price() + 2
}

type WhippedCream struct {
	coffee CoffeeItem
}

func NewWhippedCream(c CoffeeItem) *WhippedCream {
	return &WhippedCream{coffee: c}
}

func (w *WhippedCream) CoffeeItem() string {
	return w.coffee.CoffeeItem() + " + Whipped Cream"
}

func (w *WhippedCream) price() int {
	return w.coffee.price() + 5
}

type Chocolate struct {
	coffee CoffeeItem
}

func NewChocolate(c CoffeeItem) *Chocolate {
	return &Chocolate{coffee: c}
}

func (c *Chocolate) CoffeeItem() string {
	return c.coffee.CoffeeItem() + " + Chocolate"
}

func (c *Chocolate) price() int {
	return c.coffee.price() + 4
}

type Caramel struct {
	coffee CoffeeItem
}

func NewCaramel(c CoffeeItem) *Caramel {
	return &Caramel{coffee: c}
}

func (c *Caramel) CoffeeItem() string {
	return c.coffee.CoffeeItem() + " + Caramel"
}

func (c *Caramel) price() int {
	return c.coffee.price() + 3
}

type VanillaSyrup struct {
	coffee CoffeeItem
}

func NewVanillaSyrup(c CoffeeItem) *VanillaSyrup {
	return &VanillaSyrup{coffee: c}
}

func (v *VanillaSyrup) CoffeeItem() string {
	return v.coffee.CoffeeItem() + " + Vanilla Syrup"
}

func (v *VanillaSyrup) price() int {
	return v.coffee.price() + 3
}
