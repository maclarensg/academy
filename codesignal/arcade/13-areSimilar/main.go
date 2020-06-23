package main

import "fmt"

func areSimilar(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	b1 := make([]int, len(b))
	copy(b1, b)

	// Check for each array contains same element
	for _, i := range a {

		for k, j := range b1 {
			if i == j {
				// remove j at pos k of b1
				b1 = append(b1[:k], b1[k+1:]...)
				break
			}
		}

	}

	// when does not contain exact same elements in both array
	if len(b1) != 0 {
		return false
	}

	// if swap at most on pair then the matches are len(a) - 2
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			count++
		}
	}
	return count >= len(a)-2
}

func main() {
	fmt.Println(areSimilar([]int{4, 6, 3}, []int{3, 4, 6}))
}
