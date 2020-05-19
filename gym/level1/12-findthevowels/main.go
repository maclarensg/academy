package main

import (
	"fmt"
	"strings"
)

func vowels(s string) (c int) {
	c = 0
	ns := strings.ToLower(s)
	charArr := strings.Split(ns, "")
	for _, ch := range charArr {
		switch ch {
		case "a", "e", "i", "o", "u":
			c++
		}
	}
	return c
}

func main() {
	fmt.Println(vowels("Hi there!"))
	fmt.Println(vowels("Why do you ask?"))
	fmt.Println(vowels("Why?"))
}
