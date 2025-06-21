package main

import "strconv"

type Lion struct {
	AnimalBase
	runningSpeed int
}

func (l Lion) Speak() string {
	return "Roar"
}
func (l Lion) Move() string {
	return "Run at " + strconv.Itoa(l.runningSpeed) + " km/h"
}
