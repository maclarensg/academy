package main

import "fmt"

// func chunk(items []int, chunkSize int) (chunks [][]int) {
// 	// for chunkSize < len(items) {
// 	// 	items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
// 	// }

// 	// return append(chunks, items)
// }

func chunk(items []int, chunkSize int) (chunks [][]int) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func main() {

	// p chunk([1,2,3,4], 2)
	// p chunk([1,2,3,4,5], 2)
	// p chunk([1,2,3,4,5,6,7,8], 3)
	// p chunk([1,2,3,4,5], 4)
	// p chunk([1,2,3,4,5], 10)
	fmt.Println(chunk([]int{1, 2, 3, 4}, 2))
	fmt.Println(chunk([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(chunk([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3))
	fmt.Println(chunk([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(chunk([]int{1, 2, 3, 4, 5}, 10))

}
