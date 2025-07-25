package main

type ParkingLot struct {
	title  string
	levels []Level
}

func NewParkingLot(title string, levels []Level) *ParkingLot {
	return &ParkingLot{
		title:  title,
		levels: levels,
	}
}
func (p *ParkingLot) AddLevel(level Level) {
	p.levels = append(p.levels, level)
}
func (p *ParkingLot) ParkVehicle(vehicle Vehicle) ParkingSpot {
	for _, level := range p.levels {
		spot := level.AllocateSpot(vehicle)
		if spot != nil {
			return spot
		}
	}
	return nil
}
func (p *ParkingLot) UnparkVehicle(spot ParkingSpot) {
	for _, level := range p.levels {
		level.FreeSpot(spot)
	}
}
func (p *ParkingLot) GetAvailablility() map[int]int {
	availability := make(map[int]int)
	for _, level := range p.levels {
		for _, spot := range level.ParkingSpots {
			if spot.IsAvailable() {
				availability[level.LevelNumber]++
			}
		}
	}
	return availability
}
