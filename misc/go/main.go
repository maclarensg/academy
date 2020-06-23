package main

import (
	"fmt"
)

func firstNotRepeatingCharacter(s string) string {
	sa := stringSplit(s)
	csa := stringSplit(s)
	imap := make(map[string]int)
	for _, r := range sa {
		imap[r]++
		if imap[r] > 1 {
			csa = removeC(csa, r)
		}
	}

	if len(csa) > 0 {
		return csa[0]
	}

	return "_"
}

func stringSplit(s string) (sa []string) {
	for _, r := range s {
		sa = append(sa, fmt.Sprintf("%c", r))
	}
	return sa
}

func removeC(sa []string, r string) (rs []string) {
	for _, s := range sa {

		if s == r {
			continue
		} else {
			rs = append(rs, s)
		}
	}
	return rs
}

func main() {
	// fmt.Println(removeC([]string{"a", "b", "a", "c", "a", "b", "a", "d"}, "a"))
	fmt.Println(firstNotRepeatingCharacter("abacabad"))
	fmt.Println(firstNotRepeatingCharacter("abacabaabacaba"))
}
