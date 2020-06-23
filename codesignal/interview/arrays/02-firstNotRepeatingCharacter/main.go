package main

import "fmt"

func firstNotRepeatingCharacter(s string) string {
	imap := make(map[string]int)
	var seq []string
	for _, r := range s {
		if _, ok := imap[string(r)]; !ok {
			seq = append(seq, string(r))
		}
		imap[string(r)]++
	}

	for _, s := range seq {
		if imap[s] == 1 {
			return s
		}
	}
	return "_"
}

func main() {
	fmt.Println(firstNotRepeatingCharacter("aaabbbbaaaec"))
}
