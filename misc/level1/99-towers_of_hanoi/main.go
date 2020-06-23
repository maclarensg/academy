package main

import "fmt"

func move(f string, t string) {
	fmt.Printf("Move disc from %s to %s!\n", f, t)
}

func hanoi(n int, f string, h string, t string) {
	if n == 0 {
		return
	}
	hanoi(n-1, f, t, h)
	move(f, t)
	hanoi(n-1, h, f, t)
}

func main() {
	hanoi(1, "A", "B", "C")
}
