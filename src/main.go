package main

import (
	"fmt"

	"github.com/victorfernandesraton/salesman-traveling-problem/src/calculator"
	"github.com/victorfernandesraton/salesman-traveling-problem/src/route"
)

func main() {

	var cities = [][]int{
		{0, 2, 5, 7}, // a b c d
		{2, 0, 8, 3}, // b
		{5, 8, 0, 1}, // c
		{7, 3, 1, 0}} // d

	gg := route.Permute([]int{1, 2, 3})

	res := calculator.Calcule(gg, cities)

	fmt.Println()
	fmt.Println("shortestDistance:", res.Distance)
	fmt.Println("shortestPaths:", res.Path)
}
