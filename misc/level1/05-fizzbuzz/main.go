package main

import "fmt"

func fizzbuzz(n int) {
	for i := 1; i <= n; i++ {
		isMult3 := (i % 3) == 0
		isMult5 := (i % 5) == 0
		if isMult3 && isMult5 {
			fmt.Printf("FizzBuzz ")
		} else if isMult3 {
			fmt.Printf("Fizz ")
		} else if isMult5 {
			fmt.Printf("Buzz ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")
}

func main() {
	n := 100
	fizzbuzz(n)
}
