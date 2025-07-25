package main

import "math"

type FairCalculation struct {
	ratePerHour map[VehicleType]float64
}

func NewFairCalculation(ratePerHour map[VehicleType]float64) *FairCalculation {
	return &FairCalculation{
		ratePerHour: ratePerHour,
	}
}
func (f *FairCalculation) CalculateFare(vehicle Vehicle, hoursParked float64) float64 {
	rate, exists := f.ratePerHour[vehicle.GetType()]
	if !exists {
		return 0.0 // or handle error appropriately
	}
	return math.Ceil(rate * float64(hoursParked))
}
