package simplemath

import "math"

func Sqrt(value int) int {
	v := math.Sqrt(float64(value))
	return int(v)
}
