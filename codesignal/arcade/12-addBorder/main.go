package main

import "fmt"

func addBorder(picture []string) []string {
	fwidth := len(picture[0]) + 2
	for i := 0; i < len(picture); i++ {
		picture[i] = "*" + picture[i] + "*"
	}
	frame := makeNframe(fwidth)
	// add bottom frame
	picture = append(picture, frame)
	// add top frame
	picture = append([]string{frame}, picture...)

	return picture
}

func makeNframe(n int) (f string) {
	for n > 0 {
		f += "*"
		n--
	}
	return f
}

func main() {
	picture := []string{
		"abc",
		"ded",
	}
	fmt.Println(addBorder(picture))
}
