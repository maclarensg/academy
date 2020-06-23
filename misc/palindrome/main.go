package main

import (
	"fmt"
)

func palindrome(inputString string) bool {
	l := len(inputString)
	m := l / 2

	for i := 0; i < m; i++ {
		if inputString[i] != inputString[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Is Palindrome abba:", palindrome("abba"))
	fmt.Println("Is Palindrome babba:", palindrome("babba"))
}
