package main

import "fmt"

func firstDuplicate(a []int) int {
	imap := make(map[int]int)
	for _, n := range a {
		imap[n]++
		if imap[n] > 1 {
			return n
		}
	}
	return -1
}
func main() {
	fmt.Println(firstDuplicate([]int{2, 1, 3, 5, 3, 2}))
}
