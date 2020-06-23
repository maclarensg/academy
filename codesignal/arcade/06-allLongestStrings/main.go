package main

import "fmt"

func allLongestStrings(inputArray []string) []string {

	var result []string
	maxLen := 0
	for _, str := range inputArray {
		if len(str) == maxLen {
			result = append(result, str)
		}
		if len(str) > maxLen {
			maxLen = len(str)
			result = []string{str}
		}
	}

	return result
}
func main() {
	inputArray := []string{
		"abc",
		"eeee",
		"abcd",
		"dcd",
	}
	fmt.Println(allLongestStrings(inputArray))
}
