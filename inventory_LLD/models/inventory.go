package models

import (
	//"inventory_LLD/strategy"
	"sync"
)

var mutex = &sync.Mutex{}

type inventory struct {
	warehouses        []*Warehouse
	replenishstrategy ReplenishmentStrategy
}

var instance *inventory

func GetInstance() *inventory {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			instance = &inventory{warehouses: make([]*Warehouse, 0)}
		}
	}
	return instance
}

func (i *inventory) AddWarehouse(ws *Warehouse) {
	i.warehouses = append(i.warehouses, ws)
}

func (i *inventory) SetReplenishmentStrategy(rs ReplenishmentStrategy) {
	i.replenishstrategy = rs
}
