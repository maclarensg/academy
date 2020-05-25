package main

import "fmt"

func reverseInt(x int) int {
	var p int

	if x < 0 {
		p = 0 - x
	} else {
		p = x
	}

	res := 0
	for p > 0 {
		mod := p % 10
		p = p / 10
		res = 10*res + mod
	}

	if x < 0 {
		res = 0 - res
	}

	return res
}
func main() {
	fmt.Println(reverseInt(-511))
}
