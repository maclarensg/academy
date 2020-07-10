package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupingDishes(dishes [][]string) [][]string {
	iMap := make(map[string][]string)

	for _, dish := range dishes {

		for i := 1; i < len(dish); i++ {
			iMap[dish[i]] = append(iMap[dish[i]], dish[0])
		}
	}

	ingredients := [][]string{}
	for ind, dishes := range iMap {
		if len(dishes) > 1 {
			sort.Strings(dishes)
			answers := append([]string{ind}, dishes...)
			ingredients = append(ingredients, answers)
		}
	}

	sort.Slice(ingredients, func(i, j int) bool {
		return strings.Compare(ingredients[i][0], ingredients[j][0]) < 0
	})

	return ingredients
}

func main() {
	dishes := [][]string{
		{"Salad", "Tomato", "Cucumber", "Salad", "Sauce"},
		{"Pizza", "Tomato", "Sausage", "Sauce", "Dough"},
		{"Quesadilla", "Chicken", "Cheese", "Sauce"},
		{"Sandwich", "Salad", "Bread", "Tomato", "Cheese"},
	}
	// dishes := [][]string{
	// 	{"Pasta", "Tomato Sauce", "Onions", "Garlic"},
	// 	{"Chicken Curry", "Chicken", "Curry Sauce"},
	// 	{"Fried Rice", "Rice", "Onions", "Nuts"},
	// 	{"Salad", "Spinach", "Nuts"},
	// 	{"Sandwich", "Cheese", "Bread"},
	// 	{"Quesadilla", "Chicken", "Cheese"},
	// }
	// dishes := [][]string{
	// 	{"Pasta", "Tomato Sauce", "Onions", "Garlic"},
	// 	{"Chicken Curry", "Chicken", "Curry Sauce"},
	// 	{"Fried Rice", "Rice", "Onion", "Nuts"},
	// 	{"Salad", "Spinach", "Nut"},
	// 	{"Sandwich", "Cheese", "Bread"},
	// 	{"Quesadilla", "Chickens", "Cheeseeee"},
	// }

	fmt.Println(groupingDishes(dishes))
}
