package main

import "fmt"

func makeArrayConsecutive2(statues []int) int {
	sort(statues)

	p := statues[0]
	c := 0
	for i := 1; i < len(statues); i++ {
		d := statues[i] - p
		if d > 1 {
			c += d - 1
		}
		p = statues[i]
	}
	return c
}

func sort(a []int) {
	jlen := len(a) - 1
	for i := 0; i < len(a); i++ {
		for j := 0; j < jlen; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
		jlen--
	}
}

func main() {
	statues := []int{6, 2, 3, 8}
	fmt.Println(makeArrayConsecutive2(statues))
}
