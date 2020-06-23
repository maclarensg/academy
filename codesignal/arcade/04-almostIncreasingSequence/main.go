package main

import "fmt"

// 1, 2, 1, 2
// 10, 1, 2, 3
func almostIncreasingSequence(sequence []int) bool {
	c := 0
	for i := 1; i < len(sequence); i++ {
		if sequence[i] <= sequence[i-1] {
			c++
			if i+1 < len(sequence) && i-2 >= 0 {
				if sequence[i] <= sequence[i-2] && sequence[i+1] <= sequence[i-1] {
					return false
				}
			}
		}
	}
	return c <= 1
}

func main() {
	seq1 := []int{1, 2, 5, 3, 5}
	// seq2 := []int{10, 1, 2, 3}
	// seq3 := []int{10, 1, 2, 3}
	// seq4 := []int{10, 1, 2, 3}

	fmt.Println(almostIncreasingSequence(seq1))
	// fmt.Println(almostIncreasingSequence(seq2))
}
