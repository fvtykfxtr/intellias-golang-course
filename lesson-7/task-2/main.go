package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string
	var min, max int32
	numbersSlice := strings.Fields(input)

	if firstNumber, err := strconv.Atoi(numbersSlice[0]); err == nil {
		initExtremum := int32(firstNumber)
		min, max = initExtremum, initExtremum
	}

	for _, element := range numbersSlice[1:] {
		if number, err := strconv.Atoi(element); err == nil {
			potExtremum := int32(number)

			if potExtremum > max {
				max = potExtremum
			}

			if potExtremum < min {
				min = potExtremum
			}

		}
	}

	result = fmt.Sprintf("%d %d", max, min)
	fmt.Println(result)
}
