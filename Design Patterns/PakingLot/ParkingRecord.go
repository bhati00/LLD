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
func (pr *ParkingRecord) GetDuration() time.Duration {
	if pr.exitTime.IsZero() {
		return time.Since(pr.entryTime)
	}
	return pr.exitTime.Sub(pr.entryTime)
}
