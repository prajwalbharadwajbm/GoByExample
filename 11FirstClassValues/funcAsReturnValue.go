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
	// When functions are passed/returned, their environment comes with them.
	// We must understand the closure here(Fucntion+its environment), so when we pass function as argument, or return it, we also pass the closure, not just the function.
	// In this case, we see that o_x and o_y is also sent with the returned function for it to calculate once its called.
	return fn
}

func main() {
	EuclideanDist1 := CalculateEuclideanDistance(0, 0)
	EuclideanDist2 := CalculateEuclideanDistance(2, 2)
	fmt.Println(EuclideanDist1(2, 2))
	fmt.Println(EuclideanDist2(2, 2))
}
