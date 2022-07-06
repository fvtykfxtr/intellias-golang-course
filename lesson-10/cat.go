package main

import "math/rand"

type cat struct {
	animal
}

func newCat() farmAnimal {
	minWeight := 3
	maxWeight := 7
	avgWeight := rand.Intn(maxWeight-minWeight) + minWeight

	return cat{
		animal: animal{
			species:              "cat",
			weight:               avgWeight,
			foodConsumptionPerKg: 7,
		},
	}
}
