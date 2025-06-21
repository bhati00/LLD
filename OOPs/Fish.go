package main

import "strconv"

type Fish struct {
	AnimalBase
	swimmingSpeed int
}

func (f Fish) Speak() string {
	return "Blub"
}
func (f Fish) Move() string {
	return "Swim at " + strconv.Itoa(f.swimmingSpeed) + " km/h"
}
