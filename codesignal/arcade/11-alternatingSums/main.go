package main

import "fmt"

func alternatingSums(a []int) []int {
	if len(a) == 1 {
		a = append(a, 0)
		return a
	}

	if len(a) == 2 {
		return a
	}

	for i := 2; i < len(a); i++ {
		r := i % 2
		a[r] += a[i]
	}

	return a[:2]
}

func main() {
	fmt.Println(alternatingSums([]int{1}))
	fmt.Println(alternatingSums([]int{1, 2}))
	fmt.Println(alternatingSums([]int{1, 2, 3}))
	fmt.Println(alternatingSums([]int{50, 60, 60, 45, 70}))

}
