package main

import (
	"fmt"
)

func anagrams(s1 string, s2 string) bool {
	return score(s1) == score(s2)
}

func score(s string) int {
	// var arr []byte
	sBytes := []byte(s)
	score := 0
	for _, e := range sBytes {
		// downcase
		if e >= 65 && e <= 90 {
			e = e + 32
		}
		//accept alphanumeric
		if (e >= 0 && e <= 9) || (e >= 97 && e <= 122) {
			score = score + int(e)
		}
	}
	return score
}

func main() {
	fmt.Println(anagrams("rail safety", "fairy tales"))
	fmt.Println(anagrams("RAIL! SAFETY!", "fairy tales"))
}
