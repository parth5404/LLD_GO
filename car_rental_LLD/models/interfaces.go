package models

type VehicleType int
type SubVehicleType int

const (
	CAR VehicleType = iota
	BIKE
	TRUCK
)
const (
	SUV SubVehicleType = iota
	SEDAN
	HATCHBACK
)
const (
	BIKE1 SubVehicleType = iota
)

type IVehicle interface {
	GetRentalPricePerDay() float64
	GetLicenceNum() string
	SetAvailability(status bool)
	IsAvailable() bool
	GetVehicleType() VehicleType
	GetSubVehicleType() SubVehicleType
}

type IPaymentProcessor interface {
	Pay(amount float64) bool
}
