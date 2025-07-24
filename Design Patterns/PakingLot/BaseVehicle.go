package main

type VehicleType int

const (
	Car VehicleType = iota
	Motorcycle
	Bus
	Truck
)

type Vehicle interface {
	GetLicensePlate() string
	GetType() VehicleType
	GetSize() SpotType
}
type BaseVehicle struct {
	LicensePlate string
	Type         VehicleType
	Size         SpotType // Size in terms of parking space required
}

func NewBaseVehicle(licensePlate string, vehicleType VehicleType, size SpotType) *BaseVehicle {
	return &BaseVehicle{
		LicensePlate: licensePlate,
		Type:         vehicleType,
		Size:         size,
	}
}
func (v *BaseVehicle) GetLicensePlate() string {
	return v.LicensePlate
}
func (v *BaseVehicle) GetType() VehicleType {
	return v.Type
}
func (v *BaseVehicle) GetSize() SpotType {
	return v.Size
}
