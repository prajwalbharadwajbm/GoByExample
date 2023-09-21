package main

import (
	"fmt"
)

type Point struct {
	x, y, z float64
}

func newPoint(x, y, z float64) Point {
	return Point{x, y, z}
}

func main() {
	point := newPoint(1, 2, 3)
	fmt.Println(point)
}

//in progress
