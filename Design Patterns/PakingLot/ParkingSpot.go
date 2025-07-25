package main

type SpotType int

const (
	Compact SpotType = iota
	Large
	Handicapped
)

type ParkingSpot interface {
	IsAvailable() bool
	GetSpotType() SpotType
	GetSpotNumber() string
	GetVehicle() Vehicle
	ParkVehicle(vehicle Vehicle) bool
	RemoveVehicle() Vehicle
	GetLevelNumber() int
}

// shared base
type BaseSpot struct {
	spotNumber  string
	levelNumber int
	isAvailable bool
	spotType    SpotType
	vehicle     Vehicle
}

func NewBaseSpot(spotType SpotType, levelNumber int, spotNumber string) *BaseSpot {
	return &BaseSpot{
		isAvailable: true,
		spotType:    spotType,
		vehicle:     nil,
		levelNumber: levelNumber,
		spotNumber:  spotNumber,
	}
}

func (bs *BaseSpot) IsAvailable() bool     { return bs.isAvailable }
func (bs *BaseSpot) GetSpotNumber() string { return bs.spotNumber }
func (bs *BaseSpot) GetSpotType() SpotType { return bs.spotType }
func (bs *BaseSpot) GetVehicle() Vehicle   { return bs.vehicle }
func (bs *BaseSpot) GetLevelNumber() int   { return bs.levelNumber }
func (bs *BaseSpot) RemoveVehicle() Vehicle {
	v := bs.vehicle
	bs.vehicle = nil
	bs.isAvailable = true
	return v
}

func (bs *BaseSpot) ParkVehicle(v Vehicle) bool {
	if !bs.isAvailable {
		return false
	}
	bs.vehicle = v
	bs.isAvailable = false
	return true
}

type CompactSpot struct {
	BaseSpot
}

type LargeSpot struct {
	BaseSpot
}

type HandicappedSpot struct {
	BaseSpot
}
