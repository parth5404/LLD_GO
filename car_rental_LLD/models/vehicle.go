package models

type Vehicle struct {
	licencePlate   string
	available      bool
	model          string
	company        string
	ratePerDay     float64
	vehicleType    VehicleType
	subVehicleType SubVehicleType
}

func NewVehicle(licencePlate string, available bool, model string,
	company string, ratePerDay float64, vehicleType VehicleType) *Vehicle {
	return &Vehicle{
		licencePlate: licencePlate,
		available:    available,
		model:        model,
		company:      company,
		ratePerDay:   ratePerDay,
		vehicleType:  vehicleType,
	}
}
func (vehicle *Vehicle) GetRentalPricePerDay() float64 {
	return vehicle.ratePerDay
}
func (vehicle *Vehicle) GetLicenceNum() string {
	return vehicle.licencePlate
}
func (vehicle *Vehicle) SetAvailability(status bool) {
	vehicle.available = status
}
func (vehicle *Vehicle) IsAvailable() bool {
	return vehicle.available
}
func (vehicle *Vehicle) GetVehicleType() VehicleType {
	return vehicle.vehicleType
}
func (vehicle *Vehicle) GetSubVehicleType() SubVehicleType {
	return vehicle.subVehicleType
}
