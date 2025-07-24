package main

type FairCalculation struct {
	ratePerHour map[VehicleType]float64
}

func NewFairCalculation(ratePerHour map[VehicleType]float64) *FairCalculation {
	return &FairCalculation{
		ratePerHour: ratePerHour,
	}
}
func (f *FairCalculation) CalculateFare(vehicle Vehicle, hoursParked int) float64 {
	rate, exists := f.ratePerHour[vehicle.GetType()]
	if !exists {
		return 0.0 // or handle error appropriately
	}
	return rate * float64(hoursParked)
}
