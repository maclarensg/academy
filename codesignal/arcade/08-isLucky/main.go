package main

import (
	"fmt"
	"math"
	"strconv"
)

func isLucky(n int) bool {
	l := intLen(n)
	mid := l / 2
	div := int(math.Pow(10, float64(mid)))
	return sumI(n/div) == sumI(n%div)
}

func sumI(i int) int {
	sum := 0
	str := fmt.Sprintf("%d", i)
	for _, r := range str {
		s := fmt.Sprintf("%c", r)
		val, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error:", err)
		}
		sum += val
	}
	return sum
}

func intLen(n int) int {
	return len(fmt.Sprintf("%d", n))
}

func main() {
	fmt.Println(isLucky(1230))
}
