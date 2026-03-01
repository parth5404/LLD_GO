package observer

import "fmt"

type EmailObserver struct {
	id   string
	name string
}

func NewEmailObserver(id string, name string) *EmailObserver {
	return &EmailObserver{
		id:   id,
		name: name,
	}
}

func (a *EmailObserver) Update() {
	fmt.Printf("Email %v (%v) is notified\n", a.id, a.name)
}
