package models

type User struct {
	id           int
	name         string
	email        string
	reservations []*Reservations
}

func NewUser(id int, name string, email string) *User {
	return &User{
		id:           id,
		name:         name,
		email:        email,
		reservations: make([]*Reservations, 0),
	}
}
