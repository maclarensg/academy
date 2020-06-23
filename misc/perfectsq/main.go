package main

import (
	"fmt"
	"math"
)

func nextPerfectSq(n int) int {
	nextN := int(math.Sqrt(float64(n))) + 1

	return nextN * nextN
}
func square(n int) bool {
	v := math.Floor(math.Sqrt(float64(n)))
	return n == int(v*v)
}

func solve(n int) int {
	ans := -1.0
	//int(math.Max(math.Sqrt(float64(n)), 5.0))
	for i := 1; i < 5; i++ {
		// fmt.Println(math.Sqrt(float64(n)), 5.0)
		r := float64(n-i*i) / float64(2*i)

		if r > 0 && r == math.Floor(r) {
			ans = math.Pow(r, 2)
		}

		fmt.Println(">", r, ans)
	}
	return int(ans)
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, ":", solve(i))
	}
}
