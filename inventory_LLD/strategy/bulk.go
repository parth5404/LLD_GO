package strategy

import (
	"fmt"
	"inventory_LLD/models"
)

type BulkReplenishment struct{}

func (b *BulkReplenishment) Replenish(product models.InventoryItem) {
	fmt.Println("Bulk Replenishment strategy")
}
