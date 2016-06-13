package main

import "math"

func Solution(X int, Y int, D int) int {
	res := math.Ceil(float64(Y-X) / float64(D))
	return int(res)
}
