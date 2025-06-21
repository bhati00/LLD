package main

import (
	"fmt"
)

func main() {
	lions := []Lion{
		{AnimalBase: AnimalBase{name: "Simba", age: 5, habitat: "Savannah"}, runningSpeed: 80},
		{AnimalBase: AnimalBase{name: "Nala", age: 4, habitat: "Savannah"}, runningSpeed: 75},
	}
	fishes := []Fish{
		{AnimalBase: AnimalBase{name: "Nemo", age: 2, habitat: "Coral Reef"}, swimmingSpeed: 10},
		{AnimalBase: AnimalBase{name: "Dory", age: 3, habitat: "Ocean"}, swimmingSpeed: 12},
	}
	zoo := []Animal{} // using the Animal interface to hold different animal types (polymorphism)
	for _, lion := range lions {
		zoo = append(zoo, lion)
	}
	for _, fish := range fishes {
		zoo = append(zoo, fish)
	}
	for _, animal := range zoo {
		fmt.Println("-----")
		fmt.Println("Name:", animal.GetName())
		fmt.Println("Habitat:", animal.GetHabitat())
		fmt.Println("Age:", animal.GetAge())
		fmt.Println("Speak:", animal.Speak())
		fmt.Println("Move:", animal.Move())
	}

}
