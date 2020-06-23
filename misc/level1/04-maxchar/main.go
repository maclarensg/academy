package main

import (
	"fmt"
	"strings"
)

func maxchar(str string) string {
	strMap := make(map[string]int)

	sArray := strings.Split(str, "")

	max := 0
	maxC := ""

	for i := 0; i < len(sArray); i++ {
		_, ok := strMap[sArray[i]]

		if ok {
			strMap[sArray[i]]++
		} else {
			strMap[sArray[i]] = 1
		}

		if strMap[sArray[i]] > max {
			max = strMap[sArray[i]]
			maxC = sArray[i]
		}
	}

	return maxC
}

func main() {
	fmt.Println(maxchar("abccccccd"))
	fmt.Println(maxchar("apple 123111"))
}
