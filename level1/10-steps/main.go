package main

import "fmt"

// func steps(n int) {
// 	for i := 0; i < n; i++ {
// 		fmt.Printf("\"")

// 		for j := 0; j < i+1; j++ {
// 			fmt.Printf("#")
// 		}

// 		for k := 0; k < n-i-1; k++ {
// 			fmt.Printf(" ")
// 		}

// 		fmt.Println("\"")
// 	}
// }

func stepsR(n int, o int) {
	if n > 0 {
		stepsR(n-1, o)
		fmt.Printf("\"")
		for i := 0; i < n; i++ {
			fmt.Printf("#")
		}
		for k := 0; k < o-n; k++ {
			fmt.Printf(" ")
		}
		fmt.Println("\"")
	} else {
		return
	}
}

func steps(n int) {
	stepsR(n, n)
}

func main() {
	steps(2)
	steps(3)
	steps(4)
}
