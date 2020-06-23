package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

/* ---------- board and methods ---------- */
type board [][]int

func (b *board) print() {
	for _, row := range *b {
		fmt.Println(row)
	}
}

// return rows of the board
func (b *board) rows() (r [][]int) {
	return [][]int(*b)
}

// return columns of the board
func (b *board) columns() (r [][]int) {
	m := [][]int(*b)
	r = [][]int{}
	for x := 0; x < 9; x++ {
		c := []int{}
		for y := 0; y < 9; y++ {
			c = append(c, m[y][x])
		}
		r = append(r, c)
	}
	return r
}

// return blocks of the board
func (b *board) blocks() (r [][]int) {
	m := [][]int(*b)

	r = [][]int{}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			b := []int{}
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					b = append(b, m[i+(y*3)][j+(x*3)])
				}
			}
			r = append(r, b)
		}
	}
	return r
}

// Checks if Board is a valide solution
func (b *board) validate() (res bool) {

	for _, e := range b.rows() {
		if !checkAllNine(e) {
			return false
		}
	}

	for _, e := range b.columns() {
		if !checkAllNine(e) {
			return false
		}
	}

	for _, e := range b.blocks() {
		if !checkAllNine(e) {
			return false
		}
	}

	return true
}

func (b *board) solve(inter bool, all bool) {
	m := [][]int(*b)
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if m[y][x] == 0 {
				for _, p := range b.possibles(y, x) {
					m[y][x] = p
					b.solve(inter, all)
					m[y][x] = 0
				}
				return
			}
		}
	}

	b.print()

	if !all {
		return
	}

	if inter {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("More? ")
		reader.ReadString('\n')
	}
}

func (b *board) possibles(y int, x int) (p []int) {

	// set are numbers that appear in either row, column or block of a pos(y,x)
	sets := []int{}

	// add numbers to sets that exists in row, column and block
	allSets := [][]int{
		b.rows()[y],
		b.columns()[x],
		b.blocks()[whichBlockYX(y, x)],
	}

	for _, set := range allSets {
		insertToSetIfNotExist(set, &sets)
	}

	// from sets find the remaining possibles of pos(y,x)
	p = []int{}

	for i := 1; i < 10; i++ {
		if !contains(sets, i) {
			p = append(p, i)
		}
	}

	return p
}

/* ---------- other board used functions  ---------- */

func insertToSetIfNotExist(input []int, set *[]int) {
	for _, e := range input {
		if e != 0 && !contains(*set, e) {
			*set = append(*set, e)
		}
	}
}

func contains(input []int, n int) (res bool) {
	for _, e := range input {
		if e == n {
			return true
		}
	}

	return false
}

// check if the entry contains from 1-9
func checkAllNine(input []int) (res bool) {
	for i := 1; i < 10; i++ {
		search := false
		for _, e := range input {
			if e == i {
				search = true
			}
		}
		if !search {
			return false
		}
	}
	return true
}

// find which block no given y,x position
func whichBlockYX(y int, x int) (bn int) {
	return (y/3)*3 + (x / 3)
}

/* ---------- file operation functions  ---------- */

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

/* ---------- Yaml ---------- */
type boards []board // this is the struct of for unmarshalling yaml file from problems or puzzle

/* ---------- Main  ---------- */

func main() {

	allSolutions := flag.Bool("all", true, "Show all solutions.")
	interactive := flag.Bool("i", true, "Interactive? (-i=false|true)")
	problemFile := flag.String("psf", "problems.yml", "Yaml file containing problems to be solved")
	puzzleFile := flag.String("pzf", "puzzles.yml", "Yaml file containing filled puzzles for validation")

	flag.Parse()

	if !fileExists(*problemFile) && !fileExists(*puzzleFile) {
		fmt.Printf("At least one of the file %s or %s needs to be present\n", *problemFile, *puzzleFile)
		os.Exit(1)
	}

	if fileExists(*problemFile) {
		fmt.Printf("Solving all problem in file %s\n", *problemFile)
		yamlFile, err := ioutil.ReadFile(*problemFile)

		if err != nil {
			fmt.Printf("Error reading YAML file: %s\n", err)
			return
		}
		var problems boards
		err = yaml.Unmarshal(yamlFile, &problems)
		if err != nil {
			fmt.Printf("Error Unmarshalling yaml data file: %s\n", err)
			return
		}

		for _, problem := range problems {
			fmt.Println("Problem:")
			problem.print()
			fmt.Println("------")
			problem.solve(*interactive, *allSolutions)
			fmt.Println("<<< end >>>")
		}
	}

	if fileExists(*puzzleFile) {
		fmt.Printf("Validating all solutions in file %s\n", *puzzleFile)
		yamlFile, err := ioutil.ReadFile(*puzzleFile)
		if err != nil {
			fmt.Printf("Error reading YAML file: %s\n", err)
			return
		}
		var solutions boards
		err = yaml.Unmarshal(yamlFile, &solutions)
		if err != nil {
			fmt.Printf("Error Unmarshalling yaml data file: %s\n", err)
			return
		}
		for _, solution := range solutions {
			solution.print()
			fmt.Println("Is valid solution?", solution.validate())
		}
	}

}
