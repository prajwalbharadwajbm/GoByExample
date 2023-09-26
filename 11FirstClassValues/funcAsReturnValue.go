package main

import (
	"fmt"
	"math"
)

// We have created a function that returns a function, which in turn again 2 parameters to return a float64 value.
func CalculateEuclideanDistance(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_x, 2))
	}
	return fn
}

func main() {
	EuclideanDist1 := CalculateEuclideanDistance(0, 0)
	EuclideanDist2 := CalculateEuclideanDistance(2, 2)
	fmt.Println(EuclideanDist1(2, 2))
	fmt.Println(EuclideanDist2(2, 2))
}
