package main

import "fmt"

func rotateImage(a [][]int) [][]int {
	n := len(a)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			t := a[i][j]
			a[i][j] = a[n-j-1][i]
			a[n-j-1][i] = a[n-i-1][n-j-1]
			a[n-i-1][n-j-1] = a[j][n-i-1]
			a[j][n-1-i] = t
		}
	}
	return a
}

func main() {
	// m := [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{4, 5, 6},
	// 	[]int{7, 8, 9},
	// }
	m := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}
	fmt.Println(rotateImage(m))
}

// i =0; j=0
// int temp = a[i][j];   // 1
// a[i][j] = a[n-j-1][i]; // 7
// a[n-j-1][i] = a[n-1-i][n-1-j]; // 9
// a[n-1-i][n-1-j] = a[j][n-1-i]; // 3
// a[j][n-1-i] = temp; // 1

// [[1,2,3],      [[7,2,1],
//  [4,5,6],  =>   [4,5,6],
//  [7,8,9]]       [9,8,3]]

// i =0; j=0
// int temp = a[i][j];    // 2
// a[i][j] = a[n-j-1][i]; // 7
// a[n-j-1][i] = a[n-1-i][n-1-j]; // 9
// a[n-1-i][n-1-j] = a[j][n-1-i]; // 3
// a[j][n-1-i] = temp; // 1
