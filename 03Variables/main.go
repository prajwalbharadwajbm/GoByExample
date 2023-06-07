package main

import "fmt"

func main() {
	var a = "First Variable declaration"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e) //When we are not declaraing so we will get default value of int i.e, 0
	var f string
	fmt.Println(f)

	g := "Apple"
	fmt.Println(g) //Walrus operator automatically

}
