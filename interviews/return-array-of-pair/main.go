package main

import "fmt"

// For each [1,2,3,3,2,1] and n = 3 , retun [1,3,2,2,3,1]
// where
//        |-x-| |-y-|
//       [1,2,3,3,2,1]
//
// x1,y1,x2,y2,x3,y3

// [1,2,3,3,2,1] => [1,3][x,2,3,x,2,1] =>[x,x,2,3,2,1] => [1,3,2,3,2,1]
// [1,3,2,3,2,1] => [2,2][1,3,x,3,x,1] =>[1,3,x,x,3,1] => [1,3,2,2,3,1]
// [1,3,2,2,3,1]

func returnTwoN(a []int, n int) []int {
	if n-1 == 0 {
		a = append(a[1:], a[0])
		a = append(a[1:], a[0])
		return a
	}
	a = append(a[1:], a[0])
	i := a[n-1]
	a = append(append(a[:n-1], a[n:]...), i)
	return returnTwoN(a, n-1)
}

func returnTwoN2(a []int, n int) (res []int) {
	i := 0
	s := n
	for n > 0 {
		res = append(res, a[i])
		res = append(res, a[i+s])
		n--
		i++
	}

	return res
}

// [ 2 3 5 1 4 7]
// x = 5
// y = 4
// [ 2 3 x 1 x 7]
// [ 2 3 x x 1 7]
func returnTwoN3(a []int, n int) []int {
	i := 0
	m := 0
	for i <= len(a)-2 {

		x := a[i]
		y := a[i+n-m]

		for j := i + n - m - 1; j > i; j-- {
			a[j+1] = a[j]
		}

		a[i] = x
		a[i+1] = y
		i += 2
		m++
	}
	return a
}

func main() {
	// fmt.Println(returnTwoN([]int{2, 5, 1, 3, 4, 7}, 3))
	// fmt.Println(returnTwoN([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	// fmt.Println(returnTwoN([]int{1, 1, 2, 2}, 2))

	// fmt.Println(returnTwoN2([]int{2, 5, 1, 3, 4, 7}, 3))
	// fmt.Println(returnTwoN2([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	// //
	fmt.Println(returnTwoN3([]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}, 5))
	fmt.Println(returnTwoN3([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println(returnTwoN3([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	fmt.Println(returnTwoN3([]int{1, 1, 2, 2}, 2))
	fmt.Println(returnTwoN3([]int{1, 2, 3, 4, 5, 6, 6, 5, 4, 3, 2, 1}, 6))
}

// [1, 2, 3, 4, 4, 3, 2, 1]
// x=1, y=4
// [x, 2, 3, 4, x, 3, 2, 1]
// [x, x, 2, 3, 4, 3, 2, 1]
// [1, 4, 2, 3, 4, 3, 2, 1]
// ------------------------
// [1, 4, 2, 3, 4, 3, 2, 1]
// x=2, y=3
// [1, 4, x, 3, 4, x, 2, 1]
// [1, 4, x, x, 3, 4, 2, 1]
// [1, 4, 2, 3, 3, 4, 2, 1]
// ------------------------
// [1, 4, 2, 3, 3, 4, 2, 1]
// x=3, y=2
// [1, 4, 2, 3, x, 4, x, 1]
// [1, 4, 2, 3, x, x, 4, 1]
// [1, 4, 2, 3, 3, 2, 4, 1]
