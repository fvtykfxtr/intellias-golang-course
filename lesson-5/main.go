package main

import (
	"fmt"
	"math"
)

const (
	applePrice = 5.99
	pearPrice  = 7.
	money      = 23.
)

func main() {
	apples := 9.
	pears := 8.
	spending := (apples * applePrice) + (pears * pearPrice)
	fmt.Printf("Скільки грошей треба витратити, щоб купити %g яблук та %g груш? %g\n", apples, pears, spending)

	pearsPurchaseLimit, _ := math.Modf(money / pearPrice)
	fmt.Println("Скільки груш ми можемо купити?", pearsPurchaseLimit)

	applesPurchaseLimit, _ := math.Modf(money / applePrice)
	fmt.Println("Скільки яблук ми можемо купити?", applesPurchaseLimit)

	applesAssumedQuantity := 2.
	pearsAssumedQuantity := 2.
	assumedQuantityCharge := (applesAssumedQuantity * applePrice) + (pearsAssumedQuantity * pearPrice)
	isPurchasePossible := assumedQuantityCharge <= money
	fmt.Printf("Чи ми можемо купити %g груші та %g яблука? %t\n", pearsAssumedQuantity, applesAssumedQuantity, isPurchasePossible)
}
