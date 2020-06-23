package main

import "fmt"

func commonCharacterCount(s1 string, s2 string) int {
	var f1, f2 string

	f1 = s1
	f2 = s2

	count := 0

	for i := 0; i < len(f1); i++ {
		for j := 0; j < len(f2); j++ {
			if s1[i] == f2[j] {
				left, _, right := f2[:j], f2[j], f2[j+1:]
				f2 = left + right
				count++
				break
			}
		}
	}

	return count
}

func main() {
	fmt.Println(commonCharacterCount("aabcc", "adcaa"))
}
