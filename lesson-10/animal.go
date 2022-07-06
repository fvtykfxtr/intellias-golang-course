package main

import (
	"fmt"
)

var animalList = [...]string{"dog", "cat", "cow"}

type farmAnimal interface {
	calculateFoodConsumption() int
	getSummary()
}

type animal struct {
	species              string
	weight               int
	foodConsumptionPerKg int
}

func (a animal) calculateFoodConsumption() int {
	return a.weight * a.foodConsumptionPerKg
}

func (a animal) getSummary() {
	montlyFoodConsumption := a.calculateFoodConsumption()
	fmt.Printf("Farm's %v is %v kilos, monthly food consumption is %v kilos.\n", a.species, a.weight, montlyFoodConsumption)
}

func createAnimal(animalType string) (farmAnimal, error) {
	switch animalType {
	case "dog":
		return newDog(), nil
	case "cat":
		return newCat(), nil
	case "cow":
		return newCow(), nil
	default:
		return nil, fmt.Errorf("wrong animal type passed")
	}
}
