package main

import "math/rand"

type cow struct {
	animal
}

func newCow() farmAnimal {
	minWeight := 350
	maxWeight := 700
	avgWeight := rand.Intn(maxWeight-minWeight) + minWeight

	return cow{
		animal: animal{
			species:              "cow",
			weight:               avgWeight,
			foodConsumptionPerKg: 25,
		},
	}
}
