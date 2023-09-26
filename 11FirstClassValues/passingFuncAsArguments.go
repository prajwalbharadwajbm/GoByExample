package main

import "fmt"

// function can take another function as argument
func doIt(argFunc func(int) int, val int) int {
	return argFunc(val)
}

func incFunc(x int) int { return x + 1 }
func decFunc(y int) int { return y - 1 }

func main() {
	//we can have named function this way, else we can also use anonymous function(which was inspired by lambda calculus)
	fmt.Println(doIt(incFunc, 1))

	//By using anonymous function
	fmt.Println(doIt(func(y int) int { return y - 1 }, 2))
}
