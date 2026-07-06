package main

import "fmt"

var addonRegistry = map[string]func(CoffeeItem) CoffeeItem{}

func RegisterAddon(
	name string,
	creator func(CoffeeItem) CoffeeItem,
) {

	addonRegistry[name] = creator

}

func GetAddon(
	name string,
) (func(CoffeeItem) CoffeeItem, error) {

	creator, ok := addonRegistry[name]

	if !ok {
		return nil, fmt.Errorf(
			"unknown addon %s",
			name,
		)
	}

	return creator, nil
}
