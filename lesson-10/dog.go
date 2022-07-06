package main

import "math/rand"

type dog struct {
	animal
}

func newDog() farmAnimal {
	minWeight := 12
	maxWeight := 25
	avgWeight := rand.Intn(maxWeight-minWeight) + minWeight

	return dog{
		animal: animal{
			species:              "dog",
			weight:               avgWeight,
			foodConsumptionPerKg: 10 / 5,
		},
	}
}
