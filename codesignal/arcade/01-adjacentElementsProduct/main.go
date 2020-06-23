package main

import "fmt"

func adjacentElementsProduct(inputArray []int) int {
	length := len(inputArray)

	max := 0
	for i := 0; i < length-1; i++ {
		p := inputArray[i] * inputArray[i+1]

		if i == 0 {
			max = p
			continue
		}

		if p > max {
			max = p
		}
	}
	return max
}

func main() {
	inputArray := []int{3, 6, -2, -5, 7, 3}
	fmt.Println(adjacentElementsProduct(inputArray))
}
