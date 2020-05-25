package main

import "fmt"

func pad(n int) (s string) {
	s = ""
	for i := n; i > 0; i-- {
		s = s + " "
	}
	return s
}

func block(n int) (s string) {
	s = ""
	for i := n; i > 0; i-- {
		s = s + "#"
	}
	return s
}

func pyramid(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(pad(n-i-1) + block(i*2+1) + pad(n-i-1))
	}
}

func main() {
	pyramid(5)
}
