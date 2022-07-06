package main

import (
	"fmt"
	"math/rand"
)

type farm struct {
	animals              []farmAnimal
	totalFoodConsumption int
}

func (f farm) getSummary() {
	fmt.Printf("Total food consumption is: %v kilos.\n", f.totalFoodConsumption)
}

func createFarm() farm {
	farmSize := rand.Intn(20) + 1
	animals := make([]farmAnimal, farmSize)
	var totalFoodConsumption int

	for i := 0; i < farmSize; i++ {
		animalType := animalList[rand.Intn(len(animalList))]
		animal, _ := createAnimal(animalType)
		totalFoodConsumption += animal.calculateFoodConsumption()
		animals[i] = animal
	}

	return farm{
		animals,
		totalFoodConsumption,
	}
}
