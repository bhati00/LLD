package main

import "time"

type ParkingRecord struct {
	level     int
	spot      SpotType
	vehicle   Vehicle
	entryTime time.Time
	exitTime  time.Time
}

func NewParkingRecord(level int, spot SpotType, vehicle Vehicle) *ParkingRecord {
	return &ParkingRecord{
		level:     level,
		spot:      spot,
		vehicle:   vehicle,
		entryTime: time.Now(),
		exitTime:  time.Time{}, // zero value, not set yet
	}
}
func (pr *ParkingRecord) SetExitTime() {
	pr.exitTime = time.Now()
}
func (pr *ParkingRecord) GetDurationHours() float64 {
	var duration time.Duration
	if pr.exitTime.IsZero() {
		duration = time.Since(pr.entryTime)
	} else {
		duration = pr.exitTime.Sub(pr.entryTime)
	}
	return duration.Hours() // returns float64
}
