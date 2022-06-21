package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int
	appeared := make(map[int]bool)

	for _, element := range arr {
		if _, ok := appeared[element]; !ok {
			result = append(result, element)
			appeared[element] = true
		}
	}

	fmt.Println(result)
}
