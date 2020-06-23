package main

import "fmt"

func reverseInParentheses(inputString string) string {
	pc := 0 // parenthesis count

	str := ""
	rstr := ""

	for i := 0; i < len(inputString); i++ {
		fmt.Println(i, ": ", pc, ",", string(inputString[i]))
		if inputString[i] == '(' {
			if pc > 0 {
				rstr += string(inputString[i])
			}
			pc++
			continue
		}
		if inputString[i] == ')' {
			pc--
		}
		if pc == 0 && len(rstr) > 0 {
			str += reverse(rstr)
			rstr = ""
			continue
		}
		if pc > 0 {
			rstr += string(inputString[i])
			continue
		}
		if string(inputString[i]) != "(" && string(inputString[i]) != ")" {
			str += string(inputString[i])
		}
	}
	return str
}
func reverse(s string) string {
	fmt.Println(s)
	rstr := ""
	s = reverseInParentheses(s)
	for i := len(s) - 1; i >= 0; i-- {
		rstr += string(s[i])
	}
	return rstr
}

func main() {
	fmt.Println("result:", reverseInParentheses("foo(bar(baz))blim"))
}

//baroof
//foorab
