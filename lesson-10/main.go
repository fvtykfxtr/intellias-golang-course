package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	farm := createFarm()

	for _, a := range farm.animals {
		a.getSummary()
	}

	farm.getSummary()
}
