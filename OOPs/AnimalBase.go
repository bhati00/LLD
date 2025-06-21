package main

// protecting fields with lowercase names, while providing access through methods (Encapsulation)
type AnimalBase struct {
	name    string
	age     int
	habitat string
}

func (a AnimalBase) GetName() string    { return a.name }
func (a AnimalBase) GetHabitat() string { return a.habitat }
func (a AnimalBase) GetAge() int        { return a.age }
