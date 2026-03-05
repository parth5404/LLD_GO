package models

type ReservationM struct {
	vehicles     map[string]IVehicle
	customer     map[string]*User
	reservations map[string]*Reservations
}

var res_id int = 0

func (rm *ReservationM) BookReservation(vtype VehicleType,
	subvtype SubVehicleType, user User, startday int, endday int) *Reservations {
	for _, v := range rm.vehicles {
		if v.IsAvailable() && v.GetVehicleType() == vtype && v.GetSubVehicleType() == subvtype {
			v.SetAvailability(false)
			res := &Reservations{
				id:       res_id,
				user:     user,
				startday: startday,
				endday:   endday,
				vehicle:  v,
				status:   ReservationStatus(CREATED),
			}
			res_id++
			user.reservations = append(user.reservations, res)
			return res
		}
	}
	return nil
}
