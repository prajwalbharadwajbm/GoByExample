package main

import "fmt"

func main() {
	var a = "First Variable declaration"
	fmt.Println(a) //Standard varible declaration

	var z string = "Better to provide type while declaring variable"
	fmt.Println(z) //Go does automatically detects the type but its better to provide type while scanning.

	var b, c = 1, 2
	fmt.Println(b, c) //Can declare multiple variable at once

	var x, y = 1, "a"
	fmt.Println(x, y) //Can declare multiple variable at once with different variable type too

	var d = true
	fmt.Println(d) //Boolean declaration

	var e int
	fmt.Println(e) //When we are not declaraing so we will get default value of int i.e, 0
	var f string
	fmt.Println(f) //default value would be an empty string ""

	g := "Apple"
	fmt.Println(g) //The :=(walrus) syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.

}
