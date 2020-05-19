package main

import "fmt"

// S export
type S = [][]string

var (
	// I export
	I = "|"
	// U export
	U = "_"
	// O export
	O S
)

// B export
func B(c int) S {
	o := make(S, c+1)
	for x := range o {
		o[x] = make([]string, c*2+1)
		for y := range o[x] {
			o[x][y] = " "
		}
	}
	return o
}

// P export
func P(p S, r, c int) {
	l := len(O)
	r *= l / 2
	c *= l
	for x, R := range p {
		for y, C := range R {
			O[x+r][y+c] = C
		}
	}
}

// R export
func R(f S) (o S) {
	c := len(f) - 1
	o = B(c)
	for x, R := range f {
		for y, C := range R {
			if y%2 < 1 && C == I {
				o[y/2][c*2+1-x*2] = U
			}
			if y%2 > 0 && C == U {
				o[y/2+1][c*2-x*2] = I
			}
		}
	}
	return
}

func fractal(n int) (o S) {
	c := 1<<n - 1
	o = B(c)

	if c < 2 {
		o[0][1] = U
		o[1][1] = U
		o[1][2] = I
		return
	}

	a := fractal(n - 1)
	b := R(a)
	w := R(b)
	v := R(w)
	z := c/2 + 1
	O = o
	P(a, 0, 0)
	if n%2 < 1 {
		P(w, 0, 1)
		P(b, 1, 0)
		P(b, 1, 1)
		O[z][0] = I
		O[z][c*2] = I
		O[z][c] = U
	} else {
		P(v, 0, 1)
		P(w, 1, 0)
		P(v, 1, 1)
		O[0][c] = U
		O[c][c] = U
		O[z][c+1] = I
	}
	for _, R := range o {
		for y, C := range R[2:] {
			if C == U && R[y] == U {
				R[y+1] = U
			}
		}
	}
	return
}

func main() {
	arr := fractal(5)
	for _, y := range arr {
		for _, x := range y {
			fmt.Printf(x)
		}
		fmt.Println()
	}
}
