package observer

import "fmt"

type AdminObserver struct {
	id   int
	name string
}

func NewAdminObserver(id int, name string) *AdminObserver {
	return &AdminObserver{
		id:   id,
		name: name,
	}
}

func (a *AdminObserver) Update() {
	fmt.Printf("Admin %v (%v) is notified\n", a.id, a.name)
}
