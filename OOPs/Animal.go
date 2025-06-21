package main

type Animal interface {
	Speak() string
	Move() string
	GetName() string
	GetHabitat() string
	GetAge() int
}
