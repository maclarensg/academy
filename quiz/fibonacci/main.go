package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func fibI(n int64) big.Int {
	var i *big.Int
	var j *big.Int

	i = big.NewInt(0)
	j = big.NewInt(1)

	if n < 2 {

		return *big.NewInt(n)
	}

	var k int64
	var l *big.Int
	for k = 2; k <= n; k++ {
		l = big.NewInt(0)
		l = l.Add(i, j)
		i, j = j, l
	}

	return *j
}

func fibR(n int64) *big.Int {

	if n < 2 {
		return big.NewInt(n)
	}
	result := big.NewInt(0)
	return result.Add(fibR(n-2), fibR(n-1))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], " <n - integer>")
		os.Exit(1)
	}

	t, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(os.Args[1], "is not a integer")
		os.Exit(1)
	}
	n := int64(t)

	// result := fibI(n)
	result := fibR(n)
	fmt.Println(result.String())

}
