package main

func shapeArea(n int) int {
	if n == 1 {
		return n
	}

	return (n+n-1)*(n+n-1) - (dsum(n-1) * 4)
}

func dsum(i int) int {
	sum := 0
	for i > 0 {
		sum += i
		i--
	}
	return sum
}

func main() {

}

// n = 1, e = 1

// n = 2, e = 5  => (2 + 1) * (2 + 1) - (1 + 0) * 4

// n = 3, e = 13 => (3 + 2) * (3 + 2) - (2 + 1) * 4

// n = 4, e = 25 => (4 + 3) * (4 + 3) - (3 + 2 + 1) * 4
