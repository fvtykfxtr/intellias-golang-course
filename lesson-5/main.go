package main

import (
	"fmt"
	"math"
)

var (
	applePrice = 5.99
	pearPrice  = 7.
	money      = 23.
)

func main() {
	apples := 9
	pears := 8
	spending := (float64(apples) * applePrice) + (float64(pears) * pearPrice)
	fmt.Printf("Скільки грошей треба витратити, щоб купити %d яблук та %d груш? %g\n", apples, pears, spending)

	pearsPurchaseLimit, _ := math.Modf(money / pearPrice)
	fmt.Println("Скільки груш ми можемо купити?", pearsPurchaseLimit)

	applesPurchaseLimit, _ := math.Modf(money / applePrice)
	fmt.Println("Скільки яблук ми можемо купити?", applesPurchaseLimit)

	applesAssumedQuantity := 2
	pearsAssumedQuantity := 2
	assumedQuantityCharge := (float64(applesAssumedQuantity) * applePrice) + (float64(pearsAssumedQuantity) * pearPrice)
	isPurchasePossible := assumedQuantityCharge <= money
	fmt.Printf("Чи ми можемо купити %d груші та %d яблука? %t\n", pearsAssumedQuantity, applesAssumedQuantity, isPurchasePossible)
}
