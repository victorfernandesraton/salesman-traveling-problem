package calculator

import (
	"fmt"

	"github.com/victorfernandesraton/salesman-traveling-problem/src/utils"
)

// Result is a struct to compose result for calculation unity
type Result struct {
	Path     string
	Distance int
}

// CalcOperatuin is a function to operate the calculos about algorith
func CalcOperatuin(elem []int, cities [][]int, done chan Result) {

	total := 0
	lastCity := 0
	for _, city := range elem {

		total += cities[lastCity][city]

		lastCity = city
	}

	total += cities[lastCity][0]

	fmt.Println("total Distance:", total)
	done <- Result{
		Distance: total,
		Path:     utils.RoutesToStr(elem),
	}
}

// Calcule has a function to calcule Result tahn more faster algoritm
func Calcule(gg, cities [][]int) *Result {
	var acc, bfr int
	var shorter *Result
	list := &[]Result{}
	done := make(chan Result)
	for _, elem := range gg {
		fmt.Println("route:", utils.RoutesToStr(elem))
		go CalcOperatuin(elem, cities, done)
		*list = append(*list, <-done)
	}

	for _, elem := range *list {
		bfr = acc
		acc = elem.Distance
		if acc < bfr {
			shorter = &Result{
				Path:     elem.Path,
				Distance: elem.Distance,
			}
		}
	}
	return shorter
}
