package models

type Car struct {
	*Vehicle
}

func NewCar(licencePlate string, available bool, model string,
	company string, ratePerDay float64, subVehicleType SubVehicleType) *Car {
	return &Car{
		&Vehicle{
			licencePlate:   licencePlate,
			available:      available,
			model:          model,
			company:        company,
			ratePerDay:     ratePerDay,
			vehicleType:    CAR,
			subVehicleType: subVehicleType,
		},
	}
}
