package main

import (
	"strconv"
)

type Level struct {
	LevelNumber  int
	ParkingSpots []ParkingSpot
}

func NewLevel(levelNumber int, large int, compact int, handicapped int) *Level {
	parkingSpots := make([]ParkingSpot, 0)
	for range large {
		parkingSpots = append(parkingSpots, &LargeSpot{*NewBaseSpot(Large, levelNumber, strconv.Itoa(levelNumber)+"-L"+strconv.Itoa(len(parkingSpots)+1))})
	}
	for range compact {
		parkingSpots = append(parkingSpots, &CompactSpot{*NewBaseSpot(Compact, levelNumber, strconv.Itoa(levelNumber)+"-C"+strconv.Itoa(len(parkingSpots)+1))})
	}
	for range handicapped {
		parkingSpots = append(parkingSpots, &HandicappedSpot{*NewBaseSpot(Handicapped, levelNumber, strconv.Itoa(levelNumber)+"-H"+strconv.Itoa(len(parkingSpots)+1))})
	}

	return &Level{
		LevelNumber:  levelNumber,
		ParkingSpots: parkingSpots,
	}
}
func (l *Level) AddParkingSpot(spot ParkingSpot) {
	l.ParkingSpots = append(l.ParkingSpots, spot)
}
func (l *Level) AllocateSpot(vehicle Vehicle) ParkingSpot {
	for _, spot := range l.ParkingSpots {
		if spot.IsAvailable() && spot.GetSpotType() == vehicle.GetSize() {
			if spot.ParkVehicle(vehicle) {
				return spot
			}
		}
	}
	return nil
}
func (l *Level) FreeSpot(spot ParkingSpot) {
	for i, s := range l.ParkingSpots {
		if s == spot {
			l.ParkingSpots[i].RemoveVehicle()
			break
		}
	}
}
