package route

import (
	"fmt"
	"strconv"
)

// Permute has a funciton to make permutation paths using matrix and initial pth
func Permute(xs []int) (permuts [][]int) {
	var rc func(a []int, k int)

	rc = func(a []int, k int) {
		fmt.Println(k, a)
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)
	fmt.Println(permuts)
	return permuts
}

// RoutesToStr is a form to print in screen th routes
func RoutesToStr(arr []int) string {

	result := "(0,"

	for _, o := range arr {
		result += strconv.Itoa(o)
		result += ","
	}

	result += "0)"

	return result

}
