package main

import "fmt"

func areFollowingPatterns(strings []string, patterns []string) bool {
	pmap := make(map[string]string)
	smap := make(map[string]string)
	for i, str := range strings {
		val, ok := pmap[patterns[i]]
		if ok {
			if val != str {
				return false
			}
		} else {
			pmap[patterns[i]] = str
		}
		val, ok = smap[str]
		if ok {
			if val != patterns[i] {
				return false
			}
		} else {
			smap[str] = patterns[i]
		}
	}
	return true
}

func main() {
	strings := []string{"cat", "dog", "dog"}
	patterns := []string{"a", "b", "b"}
	fmt.Println(areFollowingPatterns(strings, patterns) == true)
	strings = []string{"cat", "dog", "doggy"}
	patterns = []string{"a", "b", "b"}
	fmt.Println(areFollowingPatterns(strings, patterns) == false)
	strings = []string{"cat", "dog", "dog"}
	patterns = []string{"a", "b", "c"}
	fmt.Println(areFollowingPatterns(strings, patterns) == false)
}
