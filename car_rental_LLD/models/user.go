package models

type User struct {
	id           int
	name         string
	email        string
	reservations []*Reservations
}
