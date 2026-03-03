package factory

import "car_rental_LLD/models"

type VehicleFactory struct{}

func (vf *VehicleFactory) Create(licencePlate string, available bool, model string,
	company string, ratePerDay float64,
	vehicleType models.VehicleType,
	subVehicleType *models.SubVehicleType) models.IVehicle {
	switch vehicleType {
	case models.CAR:
		return models.NewCar(licencePlate, available, model, company, ratePerDay, subVehicleType)
	default:
		return nil
	}
}
