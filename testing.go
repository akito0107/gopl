package gopl

import "math"

var EPSILON = 0.0001

func DiffAbs(a, b float64) float64 {
	return math.Abs(a - b)
}

func CompareFloat(a, b float64) int {
	if DiffAbs(a, b) < EPSILON {
		return 0
	}
	if a > b {
		return 1
	}
	return -1
}
