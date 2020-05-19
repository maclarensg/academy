package main

import "fmt"

// Matrix type 2d int array
type Matrix [][]int

func (m Matrix) print() {
	for i := range m {
		fmt.Println(m[i])
	}
}

func matrix(n int) (m Matrix) {
	m = make(Matrix, n)
	for i := range m {
		m[i] = make([]int, n)
	}

	// directions map - right, down, left, up
	directions := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}

	x, y := 0, 0
	d := 0

	for i := 0; i < n*n; i++ {
		lx, ly := x, y

		m[y][x] = i + 1 // set value

		y, x = shiftPos(ly, lx, directions[d][0], directions[d][1])

		// If shifting of position in matrix is
		// - exceeds right boundary
		// - exceeds bottom boundary
		// - exceeds left boundary
		// change direction from the last position and shift
		if (x > n-1 && d == 0) || (y > n-1 && d == 1) || (x < 0 && d == 2) {
			d = (d + 1) % 4 // change direction
			y, x = shiftPos(ly, lx, directions[d][0], directions[d][1])
		}

		// Check if postion already taken, if so, change direction from the last position and shift
		if m[y][x] != 0 {
			d = (d + 1) % 4 // change direction
			y, x = shiftPos(ly, lx, directions[d][0], directions[d][1])
		}
	}

	return m
}

func shiftPos(x int, y int, x1 int, y1 int) (x2 int, y2 int) {
	return x + x1, y + y1
}

func main() {

	m := matrix(10)
	m.print()
}
