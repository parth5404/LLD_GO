package main

import (
	"fmt"
	"inventory_LLD/factory"
	"inventory_LLD/models"
)

func main() {
	fmt.Print("Sher")
	productfact := factory.NewProductFactory()
	electronicsprod := productfact.Create(models.Electronics, "SKU123", "Laptop",
		32.50, "Dell", 25, 10)
	electronicsprod.GetSKU()
}
