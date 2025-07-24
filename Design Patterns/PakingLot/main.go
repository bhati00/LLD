package main

import "time"

func main() {

	parkinglot := NewParkingLot("Main Parking Lot", []Level{})
	// creating levels
	for i := 0; i < 3; i++ {
		level := NewLevel(i, 10, 20, 5) // 10 large, 20 compact, 5 handicapped spots
		parkinglot.AddLevel(*level)
	}
	// Example vehicle
	vehicle := NewBaseVehicle("ABC123", Car, Compact)
	// Check availability
	availability := parkinglot.GetAvailablility()
	println("Parking Lot Availability before parking:")
	for level, count := range availability {
		println("Level", level, "has", count, "available spots")
	}
	// Park the vehicle
	spot := parkinglot.ParkVehicle(vehicle)
	if spot != nil {
		println("Vehicle parked at level", spot.GetSpotType(), "spot number", spot.GetVehicle().GetLicensePlate())
	}
	// create the record
	record := NewParkingRecord(spot.GetLevelNumber(), spot.GetSpotType(), spot.GetVehicle())
	// generate ticket
	ticket := NewTicket(vehicle.LicensePlate+time.Now().Format("YYYYMMDDHHMMSS"), record)
	println("Ticket generated for vehicle:", ticket.ID)
	ratePerHour := map[VehicleType]float64{
		Car:        2.0,
		Motorcycle: 1.0,
		Truck:      3.0,
	}
	fairCalculation := NewFairCalculation(ratePerHour)
	// update exit time
	record.exitTime = time.Now().Add(2 * time.Hour)                                // assuming vehicle parked for 2 hours
	totalFare := fairCalculation.CalculateFare(vehicle, int(record.GetDuration())) // 5 hours parked
	parkinglot.UnparkVehicle(spot)
	// generating receipt
	receipt := NewReceipt(vehicle.LicensePlate+time.Now().Format("YYYYMMDDHHMMSS"), ticket.ID, totalFare, true)
	println("Receipt generated for vehicle:", receipt.ID, "with fare:", receipt.Fare)
	// Check availability again
	availability = parkinglot.GetAvailablility()
	println("Parking Lot Availability after parking:")
	for level, count := range availability {
		println("Level", level, "has", count, "available spots")
	} // Add more functionality as needed

}
