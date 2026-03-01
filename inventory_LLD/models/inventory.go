package models

import (
	//"inventory_LLD/strategy"
	"sync"
)

var mutex = &sync.Mutex{}

type Inventory struct {
	warehouses        []*Warehouse
	replenishstrategy ReplenishmentStrategy
	observerList      []Observer
}

var instance *Inventory

func GetInstance() *Inventory {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			instance = &Inventory{warehouses: make([]*Warehouse, 0)}
		}
	}
	return instance
}

func (i *Inventory) AddWarehouse(ws *Warehouse) {
	i.warehouses = append(i.warehouses, ws)
}

func (i *Inventory) SetReplenishmentStrategy(rs ReplenishmentStrategy) {
	i.replenishstrategy = rs
}

func (i *Inventory) RegisterObserver(observer Observer) {
	i.observerList = append(i.observerList, observer)
}

func (i *Inventory) RemoveObserver(observer Observer) {
	n := len(i.observerList)
	for j := 0; j < n; j++ {
		if i.observerList[j] == observer {
			i.observerList = append(i.observerList[:j-1], i.observerList[j-1+1:]...)
		}
	}
}

func (i *Inventory) NotifyObservers() {
	for _, value := range i.observerList {
		value.Update()
	}
}

func (i *Inventory) ReplenishGoods(product InventoryItem) {
	i.replenishstrategy.Replenish(product)
	i.NotifyObservers()
}
