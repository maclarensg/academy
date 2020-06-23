package main

import "fmt"

func capitalize(s string) string {
	sBytes := []byte(s)
	upFlag := false // flag to detect the 1st char of a word
	for i, e := range sBytes {
		// Upcase on the first char
		if (e >= 97 && e <= 122) && upFlag == false {
			sBytes[i] = e - 32
			upFlag = true
		} else if e == 32 { //reset upcase flag so a new word can upcase
			upFlag = false
		}
	}
	return string(sBytes)
}

func main() {
	fmt.Println(capitalize("a short sentence"))
	fmt.Println(capitalize("a lazy fox"))
	fmt.Println(capitalize("look, it is working!"))
}
