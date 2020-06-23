package main

import "fmt"

func solution(str, ending string) bool {
	if len(ending) > len(str) {
		return false
	}
	trail := str[len(str)-len(ending):]

	return trail == ending
}

func main() {
	fmt.Println(solution("banana", "d"))
}
