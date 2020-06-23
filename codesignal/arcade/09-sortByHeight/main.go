package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push to heap
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop from heap
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func sortByHeight(a []int) []int {
	h := &IntHeap{}
	heap.Init(h)

	// First Pass
	for _, n := range a {
		if n == -1 {
			continue
		}
		heap.Push(h, n)
	}
	// Second Pass
	for i, n := range a {
		if n == -1 {
			continue
		}
		a[i] = (heap.Pop(h)).(int)
	}

	return a
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	a := []int{-1, 150, 190, 170, -1, -1, 160, 180}
	fmt.Println(sortByHeight(a))

}
