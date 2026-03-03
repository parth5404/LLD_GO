package models

type ReservationStatus int

const (
	PAID ReservationStatus = iota
	CREATED
	PENDING
)

type Reservations struct {
	id       int
	user     User
	startday int
	endday   int
	vehicle  IVehicle
	status   ReservationStatus
}

func NewReservation(id int, user User, startday int, endday int, vehicle IVehicle,
	status ReservationStatus) *Reservations {
	return &Reservations{
		id:       id,
		user:     user,
		startday: startday,
		endday:   endday,
		vehicle:  vehicle,
		status:   status,
	}
}

func (rs *Reservations) calculatePayment() float64 {
	return rs.vehicle.GetRentalPricePerDay() * float64(rs.endday-rs.startday)
}
