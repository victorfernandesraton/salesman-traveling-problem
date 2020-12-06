package utils

import "strconv"

func rangeSlice(start, stop int) []int {
	if start > stop {
		panic("Slice ends before it started")
	}
	xs := make([]int, stop-start)
	for i := 0; i < len(xs); i++ {
		xs[i] = i + 1 + start
	}
	return xs
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
