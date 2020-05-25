package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	sArray := strings.Split(s, "")
	len := len(s)
	mid := len / 2
	for i := 0; i < mid; i++ {
		sArray[i], sArray[len-1-i] = sArray[len-1-i], sArray[i]
	}

	return strings.Join(sArray, "")
}

func palindrome(s string) bool {
	return s == reverse(s)
}

func main() {
	fmt.Println("Is Palindrome abba:", palindrome("abba"))
	fmt.Println("Is Palindrome babba:", palindrome("babba"))
}
