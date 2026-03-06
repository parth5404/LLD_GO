package models

type ReservationM struct {
	vehicles     map[string]IVehicle
	customer     map[string]*User
	reservations map[int]*Reservations
}

var res_id int = 0

func NewReservationManager() *ReservationM {
	return &ReservationM{
		vehicles:     make(map[string]IVehicle),
		customer:     make(map[string]*User),
		reservations: make(map[int]*Reservations),
	}
}

func (rm *ReservationM) AddVehicle(v IVehicle) {
	rm.vehicles[v.GetLicenceNum()] = v
}

func (rm *ReservationM) AddUser(u *User) {
	rm.customer[u.email] = u
}

func (rm *ReservationM) BookReservation(vtype VehicleType,
	subvtype SubVehicleType, user *User, startday int, endday int) *Reservations {
	for _, v := range rm.vehicles {
		if v.IsAvailable() && v.GetVehicleType() == vtype && v.GetSubVehicleType() == subvtype {
			v.SetAvailability(false)

			res_id++ // Increment first so ID is correct system-wide

			res := &Reservations{
				id:       res_id,
				user:     user,
				startday: startday,
				endday:   endday,
				vehicle:  v,
				status:   ReservationStatus(CREATED),
			}
			user.reservations = append(user.reservations, res)
			rm.reservations[res_id] = res
			return res
		}
	}
	return nil
}

func (rm *ReservationM) Payprice(res_id int, paymentstartegy IPaymentProcessor) bool {
	reservation := rm.reservations[res_id]
	amt := reservation.calculatePayment()
	if paymentstartegy.Pay(amt) {
		reservation.status = ReservationStatus(PAID)
		reservation.vehicle.SetAvailability(true)
		return true
	}
	return false
}
