package main

type Level struct {
	LevelNumber  int
	ParkingSpots []ParkingSpot
}

func NewLevel(levelNumber int, large int, compact int, handicapped int) *Level {
	parkingSpots := make([]ParkingSpot, large+compact+handicapped)
	for range large {
		parkingSpots = append(parkingSpots, &LargeSpot{*NewBaseSpot(Large, levelNumber)})
	}
	for range compact {
		parkingSpots = append(parkingSpots, &CompactSpot{*NewBaseSpot(Compact, levelNumber)})
	}
	for range handicapped {
		parkingSpots = append(parkingSpots, &HandicappedSpot{*NewBaseSpot(Handicapped, levelNumber)})
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
