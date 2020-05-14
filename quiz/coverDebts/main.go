package main

import (
	"fmt"
	"sort"
)

func coverDebts(s float64, debts []float64, interest []float64) float64 {
	stack := getIndexes(interest)
	is := floatSlice(interest)
	ds := floatSlice(debts)

	spend := 0.0 // Set spend to 0
	for ds.sum() > 0 {
		budget := s / 10.0
		for _, i := range stack {
			if budget >= ds[i] {
				budget = budget - ds[i]
				spend = spend + ds[i]
				ds[i] = 0
			}
			if ds[i] > budget {
				ds[i] = ds[i] - budget
				spend += budget
				budget = 0
			}
		}
		ds.computeInterest(is)
	}
	return spend
}

func main() {
	fmt.Println(coverDebts(50, []float64{2, 2, 5}, []float64{200, 100, 150}))
}

func getIndexes(slice []float64) []int {
	r := []int{}
	fs := floatSlice(slice)
	tmp := make([]float64, len(slice))
	copy(tmp, slice)
	ts := floatSlice(tmp)
	for _, v := range ts.sort() {
		//fmt.Println(fs.getPos(v))
		r = append(r, int(fs.getPos(v)))
	}
	return r
}

type floatSlice []float64

func (array floatSlice) sort() []float64 {
	sort.Sort(sort.Reverse(sort.Float64Slice(array)))
	return array
}

func (array floatSlice) getPos(value float64) int32 {
	for i, v := range array {
		if v == value {
			return int32(i)
		}
	}
	return -1
}

func (array floatSlice) sum() float64 {
	total := 0.0
	for _, v := range array {
		total = total + v
	}
	return total
}

func (array floatSlice) computeInterest(interest floatSlice) {
	for i, v := range array {
		array[i] = v + v*interest[i]/100.0
	}
}
