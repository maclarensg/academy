package main

import "fmt"

func row(grid [][]string, r int) (row []string) {
	return grid[r]
}

func column(grid [][]string, c int) (col []string) {
	if c >= 0 && c < len(grid) {
		for i := 0; i < len(grid); i++ {
			col = append(col, grid[i][c])
		}
	}
	return col
}

func block(grid [][]string, b int) (col []string) {
	blks := [][]string{}
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			blkr := []string{}
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					blkr = append(blkr, grid[i+(y*3)][j+(x*3)])
				}
			}
			blks = append(blks, blkr)
		}
	}
	return blks[b]
}

func uniq(e []string) bool {
	imap := make(map[string]int)
	for _, el := range e {
		if el != "." {
			imap[el]++
			if imap[el] > 1 {
				return false
			}
		}
	}
	return true
}

func sudoku2(grid [][]string) bool {
	size := len(grid)
	for i := 0; i < size; i++ {
		if !uniq(row(grid, i)) || !uniq(column(grid, i)) || !uniq(block(grid, i)) {
			return false
		}
	}
	return true
}

func main() {
	// g := [][]string{
	// 	{".", ".", ".", "1", "4", ".", ".", "2", "."},
	// 	{".", ".", "6", ".", ".", ".", ".", ".", "."},
	// 	{".", ".", ".", ".", ".", ".", ".", ".", "."},
	// 	{".", ".", "1", ".", ".", ".", ".", ".", "."},
	// 	{".", "6", "7", ".", ".", ".", ".", ".", "9"},
	// 	{".", ".", ".", ".", ".", ".", "8", "1", "."},
	// 	{".", "3", ".", ".", ".", ".", ".", ".", "6"},
	// 	{".", ".", ".", ".", ".", "7", ".", ".", "."},
	// 	{".", ".", ".", "5", ".", ".", ".", "7", "."},
	// }
	g := [][]string{
		{".", ".", ".", ".", "2", ".", ".", "9", "."},
		{".", ".", ".", ".", "6", ".", ".", ".", "."},
		{"7", "1", ".", ".", "7", "5", ".", ".", "."},
		{".", "7", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", "8", "3", ".", ".", "."},
		{".", ".", "8", ".", ".", "7", ".", "6", "."},
		{".", ".", ".", ".", ".", "2", ".", ".", "."},
		{".", "1", ".", "2", ".", ".", ".", ".", "."},
		{".", "2", ".", ".", "3", ".", ".", ".", "."},
	}

	// fmt.Println(row(g, 0))
	// fmt.Println(column(g, 0))
	// fmt.Println(block(g, 0))
	fmt.Println(sudoku2(g))
}
