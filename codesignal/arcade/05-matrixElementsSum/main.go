package main

import "fmt"

func matrixElementsSum(matrix [][]int) (count int) {
	yLen := len(matrix)
	xLen := len(matrix[0])

	blacklist := make(map[[2]int]bool)

	for j := 0; j < yLen; j++ {
		for i := 0; i < xLen; i++ {
			if matrix[j][i] == 0 {
				// blacklist box below
				for j0 := j + 1; j0 < yLen; j0++ {
					blacklist[[2]int{j0, i}] = true
				}
			} else {
				if _, black := blacklist[[2]int{j, i}]; !black {
					count += matrix[j][i]
				}
			}
		}
	}
	return count
}

func main() {
	matrix := [][]int{
		{1, 1, 1, 0},
		{0, 5, 0, 1},
		{2, 1, 3, 10},
	}
	fmt.Println(matrixElementsSum(matrix))
}
