package strategy

import (
	"fmt"
	"inventory_LLD/models"
)

type JustInTime struct{}

func (j *JustInTime) Replenish(product models.InventoryItem) {
	fmt.Println("Bulk Replenishment strategy")
}
