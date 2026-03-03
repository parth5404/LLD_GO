package models

type Bike struct {
	*Vehicle
	subVehicleType *SubVehicleType
}

func NewBike(licencePlate string, available bool, model string,
	company string, ratePerDay float64, subVehicleType *SubVehicleType) *Bike {
	return &Bike{
		&Vehicle{
			licencePlate: licencePlate,
			available:    available,
			model:        model,
			company:      company,
			ratePerDay:   ratePerDay,
			vehicleType:  BIKE,
		},
		subVehicleType,
	}
}
