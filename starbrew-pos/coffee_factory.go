package main

import "fmt"

type CoffeeItem interface {
	CoffeeItem() string
	price() int
}
type Espresso struct {
}

func NewEspresso() *Espresso {
	return &Espresso{}
}

func (e *Espresso) CoffeeItem() string {
	return "ESPRESSO"
}

func (e *Espresso) price() int {
	return 10
}

type Latte struct {
}

func NewLatte() *Latte {
	return &Latte{}
}

func (l *Latte) CoffeeItem() string {
	return "LATTE"
}

func (l *Latte) price() int {
	return 15
}

type Cappuccino struct {
}

func NewCappuccino() *Cappuccino {
	return &Cappuccino{}
}

func (c *Cappuccino) CoffeeItem() string {
	return "CAPPUCCINO"
}

func (c *Cappuccino) price() int {
	return 20
}

type CoffeeFactory struct {
}

func NewCoffeeFactory() *CoffeeFactory {
	return &CoffeeFactory{}
}

func (cf *CoffeeFactory) GetBaseCoffee(name string) (CoffeeItem, error) {
	switch name {
	case "ESPRESSO":
		return NewEspresso(), nil
	case "LATTE":
		return NewLatte(), nil
	case "CAPPUCCINO":
		return NewCappuccino(), nil
	default:
		return nil, fmt.Errorf("unknown base coffee: %s", name)
	}
}

//	func (cf *CoffeeFactory) AddExtra(extraName string, base CoffeeItem) (CoffeeItem, error) {
//		if base == nil {
//			return nil, fmt.Errorf("base coffee cannot be nil")
//		}
//		switch extraName {
//		case "MILK":
//			return NewMilk(base), nil
//		case "WHIPPED_CREAM":
//			return NewWhippedCream(base), nil
//		case "CHOCOLATE":
//			return NewChocolate(base), nil
//		case "CARAMEL":
//			return NewCaramel(base), nil
//		case "VANILLA_SYRUP":
//			return NewVanillaSyrup(base), nil
//		default:
//			return base, fmt.Errorf("unknown extra: %s", extraName)
//		}
//	}
func ApplyAddon(coffee CoffeeItem, addon string) (CoffeeItem, error) {
	creator, err := GetAddon(addon)
	if err != nil {
		return nil, err
	}
	return creator(coffee), nil
}

func init() {
	RegisterAddon(
		"milk",
		func(c CoffeeItem) CoffeeItem {

			return NewMilk(c)

		},
	)
	RegisterAddon(
		"caramel",
		func(c CoffeeItem) CoffeeItem {

			return NewCaramel(c)

		},
	)
}
