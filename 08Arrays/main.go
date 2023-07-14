package main

import (
	"fmt"
)

func main() {
	//Lets create an array of length 5 with type integer.
	//Until we initialize value for each of those 5 position, by default it would be 0.
	var a [3]int

	fmt.Println("Before initialization")
	fmt.Println(a)

	fmt.Println("After initialization")
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)

	//find length of an array using builtin len method.
	fmt.Println("Length of array a is:", len(a))

	//to declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoDimensionalArrayEg [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimensionalArrayEg[i][j] = i + j
		}
	}
	fmt.Println(twoDimensionalArrayEg)

}
