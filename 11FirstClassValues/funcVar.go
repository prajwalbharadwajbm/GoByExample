package main

import "fmt"

// function declaration as a variable
// SYNTAX : var <func-name> func(<datatype of the parameter/parameters>) <return DataType>
var funcVar func(int) int

// We are writing this function to define as value in main
func incFn(x int) int {
	return x + 1
}

func main() {
	//We defined the value for the function variable as incFn
	funcVar = incFn
	//Now it doesn't matter if its funcVar(1) or incFn(1), it calls incFn only.
	fmt.Println(funcVar(1))
}
