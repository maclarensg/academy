package main

import (
	"fmt"
	"strconv"
)

func isCryptSolution(crypt []string, solution [][]string) bool {
	array := make([]int, 3)

	for i, c := range crypt {
		a := encrypt(c, solution)

		if len(a) > 1 && a[0] == '0' {
			return false
		}

		aint, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println(err)
			return false
		}
		array[i] = aint
	}

	return array[0]+array[1] == array[2]
}

func encode(s string, solution [][]string) string {
	for _, t := range solution {
		if t[0] == s {
			return t[1]
		}
	}
	return ""
}

func encrypt(str string, solution [][]string) string {
	res := ""
	for _, r := range str {
		c := fmt.Sprintf("%c", r)
		e := encode(c, solution)
		res += e
	}
	return res
}

func main() {
	crypt := []string{"SEND", "MORE", "MONEY"}
	solution := [][]string{
		{"O", "0"},
		{"M", "1"},
		{"Y", "2"},
		{"E", "5"},
		{"N", "6"},
		{"D", "7"},
		{"R", "8"},
		{"S", "9"},
	}
	fmt.Println(isCryptSolution(crypt, solution))
}
