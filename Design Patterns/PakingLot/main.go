package main

import "time"

func main() {

	parkinglot := NewParkingLot("Main Parking Lot", []Level{})
	for i := 0; i < 3; i++ {
		level := NewLevel(i, 10, 20, 5)
		parkinglot.AddLevel(*level)
	}
	vehicle := NewBaseVehicle("ABC123", Car, Compact)
	availability := parkinglot.GetAvailablility()
	println("Parking Lot Availability before parking:")
	for level, count := range availability {
		println("Level", level, "has", count, "available spots")
	}
	spot := parkinglot.ParkVehicle(vehicle)
	if spot != nil {
		println("Vehicle parked at level", spot.GetSpotType(), "spot number", spot.GetSpotNumber())
	}
	record := NewParkingRecord(spot.GetLevelNumber(), spot.GetSpotType(), spot.GetVehicle())
	ticket := NewTicket(vehicle.LicensePlate+time.Now().Format("YYYYMMDDHHMMSS"), record)
	println("Ticket generated for vehicle:", ticket.ID)
	availability = parkinglot.GetAvailablility()
	println("Parking Lot Availability after parking:")
	for level, count := range availability {
		println("Level", level, "has", count, "available spots")
	}
	ratePerHour := map[VehicleType]float64{
		Car:        2.0,
		Motorcycle: 1.0,
		Truck:      3.0,
	}
	fairCalculation := NewFairCalculation(ratePerHour)
	record.exitTime = time.Now().Add(2 * time.Hour)
	totalDuration := record.GetDurationHours()                         // assuming vehicle parked for 2 hours
	totalFare := fairCalculation.CalculateFare(vehicle, totalDuration) // 5 hours parked
	parkinglot.UnparkVehicle(spot)
	// generating receipt
	receipt := NewReceipt(vehicle.LicensePlate+time.Now().Format("YYYYMMDDHHMMSS"), ticket.ID, totalFare, true)
	println("Receipt generated for vehicle:", receipt.ID, "with fare:", receipt.Fare)

	// Add more functionality as needed

}
