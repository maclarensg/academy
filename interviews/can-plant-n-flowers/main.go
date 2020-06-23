package main

import "fmt"

//Plant flowers,
// if [1,0,0,0,1] and n=1 , result true
// each flower need a empty space in between
// if [1,0,0,0,1] and n=2 , result false we cannot plant more than 1 and has 1 left
// if [0, 0, 0, 0, 0] and n=3 , result true
// each flower need a empty space in between
// if [0, 0, 0, 0, 0] and n=4 , result false because has 1 left

func checkFlowerBed(fb []int, n int) bool {

	p := 0
	for _, pos := range fb {
		if n == 0 {
			break
		}

		if pos == 0 && p != 1 {
			p = 1
			n--
		} else {
			p = pos
		}

	}

	if n > 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(checkFlowerBed([]int{1, 0, 0, 0, 1}, 1) == true)
	fmt.Println(checkFlowerBed([]int{1, 0, 0, 0, 1}, 2) == false)
	fmt.Println(checkFlowerBed([]int{0, 0, 0, 0, 0}, 3) == true)
	fmt.Println(checkFlowerBed([]int{0, 0, 0, 0, 0}, 4) == false)
}
